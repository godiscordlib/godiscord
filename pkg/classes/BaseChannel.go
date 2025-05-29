package classes

import (
	types2 "godiscord.foo.ng/lib/pkg/types"
)

type BaseChannel struct {
	Base
	Type                 types2.ChannelType    `json:"type"`
	Name                 string                `json:"name"`
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
	GuildID              string                `json:"guild_id"`
	Position             int                   `json:"position"`
	Topic                string                `json:"topic"`
	Flags                int                   `json:"flags"`
	LastMessageID        string
	NSFW                 bool
	CategoryID           string
	Guild                Guild
}

type CreateChannelOptions struct {
	Name                 string                       `json:"name"`
	Type                 types2.ChannelType           `json:"type"`
	Topic                string                       `json:"topic,omitempty"`
	Bitrate              int                          `json:"bitrate,omitempty"`
	UserLimit            int                          `json:"user_limit,omitempty"`
	MessageCooldown      int                          `json:"rate_limit_per_user"`
	Position             int                          `json:"position,omitempty"`
	PermissionOverwrites []PermissionOverwrite        `json:"permission_overwrites,omitempty"`
	CategoryID           string                       `json:"parent_id,omitempty"`
	NSFW                 bool                         `json:"nsfw"`
	RTCRegion            string                       `json:"rtc_region,omitempty"`
	VideoQuality         types2.VideoQuality          `json:"video_quality_mode,omitempty"`
	DefaultReactionEmoji DefaultReaction              `json:"default_reaction_emoji"`
	AvailableTags        []ForumTag                   `json:"available_tags"`
	DefaultFormSortOrder types2.DefaultForumSortOrder `json:"default_sort_order"`
	DefaultForumLayout   types2.DefaultForumLayout    `json:"default_forum_layout"`
}

type Channel interface {
	GetType() types2.ChannelType
}

func (bd BaseChannel) GetType() types2.ChannelType {
	return bd.Type
}
