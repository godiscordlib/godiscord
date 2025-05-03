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
