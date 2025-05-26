package classes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type InviteManager struct {
}

func (im InviteManager) Get(Code string) (*Invite, error) {
	var invite_code string
	URL, err := url.Parse(Code)
	if err != nil {
		return nil, err
	}
	if URL.Scheme == "https" {
		tld := strings.Split(URL.Host, ".")[1]
		switch tld {
		case "gg":
			invite_code = strings.TrimPrefix(URL.Path, "/")
		case "com":
			path := strings.Split(URL.Path, "/")
			if path[0] != "invite" {
				return nil, fmt.Errorf("error: discord url %s is not an invite link", Code)
			}
			invite_code = strings.Split(URL.Path, "/")[2]
		default:
			return nil, fmt.Errorf("error: the code %s is not a valid invite link", Code)
		}
	} else {
		invite_code = Code
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/invites/%s", API_URL, invite_code), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var invite Invite
	json.Unmarshal(body, &invite)
	return &invite, err
}

func (im InviteManager) Delete(Code string, Reason ...string) error {
	var invite_code, reason string

	if len(Reason) > 0 {
		reason = Reason[0]
	}

	URL, err := url.Parse(Code)
	if err != nil {
		return err
	}
	if URL.Scheme == "https" {
		tld := strings.Split(URL.Host, ".")[1]
		switch tld {
		case "gg":
			invite_code = strings.TrimPrefix(URL.Path, "/")
		case "com":
			path := strings.Split(URL.Path, "/")
			if path[0] != "invite" {
				return fmt.Errorf("error: discord url %s is not an invite link", Code)
			}
			invite_code = strings.Split(URL.Path, "/")[2]
		default:
			return fmt.Errorf("error: the code %s is not a valid invite link", Code)
		}
	} else {
		invite_code = Code
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/invites/%s", API_URL, invite_code), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("X-Audit-Log-Reason", reason)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("error: unknown invite/missing permissions")
	}
	return nil
}
