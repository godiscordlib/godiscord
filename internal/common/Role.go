package common

import "github.com/AYn0nyme/godiscord/internal/enums"

type Role struct {
	Base
	Name               string           `json:"name"`
	Color              int              `json:"color"`
	Hoisted            bool             `json:"hoist"`
	IconHash           string           `json:"icon"`
	Position           int              `json:"position"`
	Permissions        enums.Permission `json:"permissions"`
	IntegrationManaged bool             `json:"managed"`
	Mentionable        bool             `json:"mentionable"`
	Flags              int              `json:"flags"`
}

type CreateRoleOptions struct {
	Name        string `json:"name,omitempty"`
	Permissions string `json:"permissions,omitempty"`
	Color       int    `json:"color,omitempty"`
	Hoist       bool   `json:"hoisted,omitempty"`
	// TODO: add icon
	UnicodeEmoji string `json:"unicode_string,omitempty"`
	Mentionable  bool   `json:"mentionable,omitempty"`
}

type EditRoleOptions struct {
	Name         string `json:"name"`
	Permissions  string `json:"permissions"`
	Color        int    `json:"color,omitempty"`
	Hoisted      bool   `json:"hoist"`
	UnicodeEmoji string `json:"unicode_string,omitempty"`
	Mentionable  bool   `json:"mentionable,omitempty"`
}
