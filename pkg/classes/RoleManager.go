package classes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RoleManager struct {
	GuildID string
}

func (rm RoleManager) GetAll() (*[]Role, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/guilds/%s/roles", API_URL, rm.GuildID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	var roles []Role
	if err = json.Unmarshal(body, &roles); err != nil {
		return nil, err
	}
	return &roles, nil
}

func (rm RoleManager) Get(RoleID string) (*Role, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/guilds/%s/roles/%s", API_URL, rm.GuildID, RoleID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	var role Role
	if err = json.Unmarshal(body, &role); err != nil {
		return nil, err
	}
	return &role, nil
}

func (rm RoleManager) Create(Options CreateRoleOptions) (*Role, error) {
	req_body, err := json.Marshal(Options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/guilds/%s/roles", API_URL, rm.GuildID), bytes.NewReader(req_body))
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
	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("%s", string(body))
	}
	var role Role
	if err = json.Unmarshal(body, &role); err != nil {
		return nil, err
	}
	return &role, nil
}

func (rm RoleManager) Edit(RoleID string, Options EditRoleOptions) (*Role, error) {
	req_body, err := json.Marshal(Options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/guilds/%s/roles/%s", API_URL, rm.GuildID, RoleID), bytes.NewReader(req_body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("X-Audit-Log-Reason", Options.Reason)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	res_body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(res_body))
	}
	var role Role
	if err = json.Unmarshal(res_body, &role); err != nil {
		return nil, err
	}
	return &role, nil
}

func (rm RoleManager) Delete(RoleID string, Reason ...string) error {
	var reason string
	if len(Reason) > 0 {
		reason = Reason[0]
	}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/guilds/%s/roles/%s", API_URL, rm.GuildID, RoleID), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("X-Audit-Log-Reason", reason)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}
	return nil
}
