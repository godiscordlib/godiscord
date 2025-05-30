package slash

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"godiscord.foo.ng/lib/pkg/classes"
)

func RegisterGuildCommands(ApplicationID string, Commands []classes.SlashCommandData, GuildID string) error {
	for _, cmd := range Commands {
		if cmd.Name == "" || cmd.Description == "" || cmd.Type == 0 {
			return errors.New("missing required fields in SlashCommandData")
		}
		reqBody, err := json.Marshal(cmd)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/applications/%s/guilds/%s/commands", classes.API_URL, ApplicationID, GuildID), bytes.NewReader(reqBody))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if res.StatusCode != 200 {
			return errors.New(string(body))
		}
	}

	return nil
}

func RegisterGlobalCommands(ApplicationID string, Commands []classes.SlashCommandData) error {
	for _, cmd := range Commands {
		if cmd.Name == "" || cmd.Description == "" || cmd.Type == 0 {
			return errors.New("missing required fields in SlashCommandData")
		}
		reqBody, err := json.Marshal(cmd)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/applications/%s/commands", classes.API_URL, ApplicationID), bytes.NewReader(reqBody))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if res.StatusCode != 200 {
			return errors.New(string(body))
		}
	}

	return nil
}
