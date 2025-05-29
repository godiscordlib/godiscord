package enums

import (
	"godiscord.foo.ng/lib/pkg/types"
)

var ChannelType = struct {
	TextChannel             types.ChannelType
	DMChannel               types.ChannelType
	VoiceChannel            types.ChannelType
	DMGroup                 types.ChannelType
	Category                types.ChannelType
	GuildAnnouncement       types.ChannelType
	GuildAnnouncementThread types.ChannelType
	PublicThread            types.ChannelType
	PrivateThread           types.ChannelType
	StageVoice              types.ChannelType
	GuildHubDirectory       types.ChannelType
	Forum                   types.ChannelType
	GuildMedia              types.ChannelType
}{
	TextChannel:             types.TextChannel,
	DMChannel:               types.DMChannel,
	VoiceChannel:            types.VoiceChannel,
	DMGroup:                 types.DMGroup,
	Category:                types.Category,
	GuildAnnouncement:       types.GuildAnnouncement,
	GuildAnnouncementThread: types.GuildAnnouncementThread,
	PublicThread:            types.PublicThread,
	PrivateThread:           types.PrivateThread,
	StageVoice:              types.StageVoice,
	GuildHubDirectory:       types.GuildHubDirectory,
	Forum:                   types.Forum,
	GuildMedia:              types.GuildMedia,
}
