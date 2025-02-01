package enums

const (
	TextChannel int = iota
	DMChannel
	VoiceChannel
	DMGroup
	Category
	GuildAnnouncement
	GuildAnnouncementThread int = iota + 4
	PublicThread
	PrivateThread
	StageVoice
	GuildHubDirectory
	Forum
	GuildMedia
)
