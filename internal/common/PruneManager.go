package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type PruneManager struct {
	Guild *Guild
}

type GetPruneCountOptions struct {
	Days  int      `json:"days"`
	Roles []string `json:"include_roles,omitempty"` // Array of role IDs to include
}
type BeginPruneOptions struct {
	Days   int      `json:"days"`
	Roles  []string `json:"include_roles,omitempty"` // Array of role IDs to include
	Reason string
}

func (pm PruneManager) GetCount(Options GetPruneCountOptions) (*int, error) {
	req_body, err := json.Marshal(Options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/guilds/%s/prune", API_URL, pm.Guild.ID), bytes.NewReader(req_body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", os.Getenv("GODISCORD_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	res_body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(res_body))
	}
	var pruned struct {
		Pruned int `json:"pruned"`
	}
	if err = json.Unmarshal(res_body, &pruned); err != nil {
		return nil, err
	}
	return &pruned.Pruned, err
}

func (pm PruneManager) Begin(Options BeginPruneOptions) error {
	req_body, err := json.Marshal(Options)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/guilds/%s/prune", API_URL, pm.Guild.ID), bytes.NewReader(req_body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("X-Audit-Log-Reason", Options.Reason)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		defer res.Body.Close()
		res_body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(res_body))
	}
	return nil
}
