package enums

import (
	"godiscord.foo.ng/lib/pkg/types"
)

var GatewayIntent = struct {
	Guilds                  types.GatewayIntent
	GuildMembers            types.GatewayIntent
	GuildModeration         types.GatewayIntent
	GuildExpressions        types.GatewayIntent
	GuildIntegrations       types.GatewayIntent
	GuildWebhooks           types.GatewayIntent
	GuildInvites            types.GatewayIntent
	GuildVoiceStates        types.GatewayIntent
	GuildPresences          types.GatewayIntent
	GuildMessages           types.GatewayIntent
	GuildMessageReactions   types.GatewayIntent
	GuildMessageTyping      types.GatewayIntent
	DirectMessages          types.GatewayIntent
	DirectMessageReactions  types.GatewayIntent
	DirectMessageTyping     types.GatewayIntent
	MessageContent          types.GatewayIntent
	GuildScheduledEvents    types.GatewayIntent
	AutoModerationConfig    types.GatewayIntent
	AutoModerationExecution types.GatewayIntent
	GuildsMessagePolls      types.GatewayIntent
	DirectMessagesPolls     types.GatewayIntent
}{
	Guilds:                  types.GI_Guilds,
	GuildMembers:            types.GI_GuildMembers,
	GuildModeration:         types.GI_GuildModeration,
	GuildExpressions:        types.GI_GuildExpressions,
	GuildIntegrations:       types.GI_GuildIntegrations,
	GuildWebhooks:           types.GI_GuildWebhooks,
	GuildInvites:            types.GI_GuildInvites,
	GuildVoiceStates:        types.GI_GuildVoiceStates,
	GuildPresences:          types.GI_GuildPresences,
	GuildMessages:           types.GI_GuildMessages,
	GuildMessageReactions:   types.GI_GuildMessageReactions,
	GuildMessageTyping:      types.GI_GuildMessageTyping,
	DirectMessages:          types.GI_DirectMessages,
	DirectMessageReactions:  types.GI_DirectMessageReactions,
	DirectMessageTyping:     types.GI_DirectMessageTyping,
	MessageContent:          types.GI_MessageContent,
	GuildScheduledEvents:    types.GI_GuildScheduledEvents,
	AutoModerationConfig:    types.GI_AutoModerationConfiguration,
	AutoModerationExecution: types.GI_AutoModerationExecution,
	GuildsMessagePolls:      types.GI_GuildsMessagePolls,
	DirectMessagesPolls:     types.GI_DirectMessagesPolls,
}
