package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var ChannelType = struct {
	GuildText               types.ChannelType
	DMChannel               types.ChannelType
	GuildVoice              types.ChannelType
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
	types.TextChannel,
	types.DMChannel,
	types.VoiceChannel,
	types.DMGroup,
	types.Category,
	types.GuildAnnouncement,
	types.GuildAnnouncementThread,
	types.PublicThread,
	types.PrivateThread,
	types.StageVoice,
	types.GuildHubDirectory,
	types.Forum,
	types.GuildMedia,
}
