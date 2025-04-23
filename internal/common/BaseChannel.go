package common

import "github.com/AYn0nyme/godiscord/internal/enums"

type BaseChannel struct {
	Base
	Type                 enums.ChannelType     `json:"type"`
	Name                 string                `json:"name"`
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
	GuildID              string                `json:"guild_id"`
	Position             int                   `json:"position"`
	Topic                string                `json:"topic"`
	Flags                int                   `json:"flags"`
}
