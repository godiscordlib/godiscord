package classes

import "godiscord.foo.ng/lib/internal/types"

type Webhook struct {
	Base
	Type          types.WebhookType `json:"type"`
	GuildID       string            `json:"guild_id"`
	ChannelID     string            `json:"channel_id"`
	User          User              `json:"user"`
	Name          string            `json:"name"`
	AvatarHash    string            `json:"avatar"`
	Token         string            `json:"token"`
	ApplicationID string            `json:"application_id"`
	SourceGuild   Guild             `json:"source_guild"`
	SourceChannel BaseChannel       `json:"source_channel"`
	URL           string            `json:"url"`
}
