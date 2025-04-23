package common

import "github.com/AYn0nyme/godiscord/internal/enums"

type ChannelMention struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Type    enums.ChannelType `json:"type"`
	GuildID string            `json:"guild_id"`
}
