package classes

import (
	"godiscord.foo.ng/lib/pkg/types"
)

type ChannelMention struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Type    types.ChannelType `json:"type"`
	GuildID string            `json:"guild_id"`
}
