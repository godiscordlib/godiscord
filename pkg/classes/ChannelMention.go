package classes

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

type ChannelMention struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Type    types.ChannelType `json:"type"`
	GuildID string            `json:"guild_id"`
}
