package common

import "github.com/AYn0nyme/godiscord/internal/enums"

type BaseChannel struct {
	Base
	Type                 enums.ChannelType     `json:"type"`
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
	Name                 string                      `json:"name"`
	Type                 enums.ChannelType           `json:"type"`
	Topic                string                      `json:"topic,omitempty"`
	Bitrate              int                         `json:"bitrate,omitempty"`
	UserLimit            int                         `json:"user_limit,omitempty"`
	MessageCooldown      int                         `json:"rate_limit_per_user"`
	Position             int                         `json:"position,omitempty"`
	PermissionOverwrites []PermissionOverwrite       `json:"permission_overwrites,omitempty"`
	CategoryID           string                      `json:"parent_id,omitempty"`
	NSFW                 bool                        `json:"nsfw"`
	RTCRegion            string                      `json:"rtc_region,omitempty"`
	VideoQuality         enums.VideoQuality          `json:"video_quality_mode,omitempty"`
	DefaultReactionEmoji DefaultReaction             `json:"default_reaction_emoji"`
	AvailableTags        []ForumTag                  `json:"available_tags"`
	DefaultFormSortOrder enums.DefaultForumSortOrder `json:"default_sort_order"`
	DefaultForumLayout   enums.DefaultForumLayout    `json:"default_forum_layout"`
}

type Channel interface {
	GetType() enums.ChannelType
}

func (bd BaseChannel) GetType() enums.ChannelType {
	return bd.Type
}
