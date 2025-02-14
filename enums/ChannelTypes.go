package enums

type ChannelType int

const (
	TextChannel ChannelType = iota
	DMChannel
	VoiceChannel
	DMGroup
	Category
	GuildAnnouncement
	GuildAnnouncementThread ChannelType = iota + 4
	PublicThread
	PrivateThread
	StageVoice
	GuildHubDirectory
	Forum
	GuildMedia
)
