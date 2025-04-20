package enums

type Permission int

const (
	CreateInstantInvite Permission = 1 << iota
	KickMembers
	BanMembers
	Administrator
	ManageChannels
	ManageGuild
	AddReactions
	ViewAuditLog
	PrioritySpeaker
	Stream
	ViewChannel
	SendMessages
	SendTTSMessages
	ManageMessages
	LinkEmbeds
	AttachFiles
	ReadMessageHistory
	MentionEveryone
	UseExternalEmojis
	ViewGuildInsights
	Connect
	Speak
	MuteMembers
	DeafenMembers
	MoveMembers
	UseVoiceActivityDetection
	ChangeNickname
	ManageNicknames
	ManageRoles
	ManageWebhooks
	ManageGuildExpressions
	UseApplicationCommands
	RequestToSpeak
	ManageEvents
	ManageThreads
	CreatePublicThread
	CreatePrivateThread
	UseExternalStickers
	SendMessagesInThreads
	UseEmbeddedActivities
	ModerateMembers
	ViewMonetizationAnalytics
	UseSoundboard
	CreateGuildExpressions
	CreateEvents
	UseExternalSounds
	SendVoiceMessages
	SendPolls
	UseExternalApps
)
