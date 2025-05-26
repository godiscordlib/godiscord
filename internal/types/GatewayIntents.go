package types

type GatewayIntent int

const (
	GI_Guilds GatewayIntent = 1 << iota
	GI_GuildMembers
	GI_GuildModeration
	GI_GuildExpressions
	GI_GuildIntegrations
	GI_GuildWebhooks
	GI_GuildInvites
	GI_GuildVoiceStates
	GI_GuildPresences
	GI_GuildMessages
	GI_GuildMessageReactions
	GI_GuildMessageTyping
	GI_DirectMessages
	GI_DirectMessageReactions
	GI_DirectMessageTyping
	GI_MessageContent
	GI_GuildScheduledEvents
)

const (
	GI_AutoModerationConfiguration GatewayIntent = 1 << 20
	GI_AutoModerationExecution     GatewayIntent = 1 << 21
	GI_GuildsMessagePolls          GatewayIntent = 1 << 24
	GI_DirectMessagesPolls         GatewayIntent = 1 << 25
)
