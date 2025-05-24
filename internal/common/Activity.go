package common

import "github.com/AYn0nyme/godiscord/internal/enums"

type Activity struct {
	Name       string              `json:"name"`
	Type       enums.ActivityType  `json:"type"`
	URL        string              `json:"url,omitempty"`
	Flags      enums.ActivityFlags `json:"flags,omitempty"` // si besoin, optionnel
	CreatedAt  int64               `json:"created_at"`
	Timestamps struct {
		Start int64 `json:"start,omitempty"`
		End   int64 `json:"end,omitempty"`
	} `json:"timestamps,omitempty"`
	ApplicationID string `json:"application_id,omitempty"`
	Details       string `json:"details,omitempty"`
	State         string `json:"state,omitempty"`
	Emoji         struct {
		Name     string `json:"name,omitempty"`
		ID       string `json:"id,omitempty"`
		Animated bool   `json:"animated,omitempty"`
	} `json:"emoji,omitempty"`
	Party struct {
		ID   string `json:"id,omitempty"`
		Size int    `json:"animated,omitempty"`
	} `json:"party,omitempty"`
	Assets struct {
	} `json:"assets,omitempty"`
	Secrets  struct{} `json:"secrets,omitempty"`
	Instance bool     `json:"instance,omitempty"`
	Buttons  []struct {
		Label string `json:"label,omitempty"`
		URL   string `json:"url,omitempty"`
	} `json:"buttons,omitempty"`
}
