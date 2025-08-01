package classes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/godiscordlib/godiscord/pkg/types"
)

type Channel struct {
	Base
	Type                 types.ChannelType     `json:"type"`
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
type ChannelInt interface {
	GetCType() types.ChannelType
}

type CreateChannelOptions struct {
	Name                 string                      `json:"name"`
	Type                 types.ChannelType           `json:"type"`
	Topic                string                      `json:"topic,omitempty"`
	Bitrate              int                         `json:"bitrate,omitempty"`
	UserLimit            int                         `json:"user_limit,omitempty"`
	MessageCooldown      int                         `json:"rate_limit_per_user"`
	Position             int                         `json:"position,omitempty"`
	PermissionOverwrites []PermissionOverwrite       `json:"permission_overwrites,omitempty"`
	CategoryID           string                      `json:"parent_id,omitempty"`
	NSFW                 bool                        `json:"nsfw"`
	RTCRegion            string                      `json:"rtc_region,omitempty"`
	VideoQuality         types.VideoQuality          `json:"video_quality_mode,omitempty"`
	DefaultReactionEmoji DefaultReaction             `json:"default_reaction_emoji"`
	AvailableTags        []ForumTag                  `json:"available_tags"`
	DefaultFormSortOrder types.DefaultForumSortOrder `json:"default_sort_order"`
	DefaultForumLayout   types.DefaultForumLayout    `json:"default_forum_layout"`
}

type EditChannelOptions struct {
	Name                 string                      `json:"name,omitempty"`
	Type                 types.ChannelType           `json:"type,omitempty"`
	Topic                string                      `json:"topic,omitempty"`
	Bitrate              int                         `json:"bitrate,omitempty"`
	UserLimit            int                         `json:"user_limit,omitempty"`
	MessageCooldown      int                         `json:"rate_limit_per_user,omitempty"`
	PermissionOverwrites []PermissionOverwrite       `json:"permission_overwrites,omitempty"`
	CategoryID           string                      `json:"parent_id,omitempty"`
	NSFW                 bool                        `json:"nsfw,omitempty"`
	Flags                types.ChannelFlag           `json:"flags,omitempty"`
	AvailableTags        []ForumTag                  `json:"available_tags,omitempty"`
	DefaultReactionEmoji DefaultReaction             `json:"default_reaction_emoji,omitempty"`
	DefaultSortOrder     types.DefaultForumSortOrder `json:"default_sort_order,omitempty"`
	DefaultForumLayout   types.DefaultForumLayout    `json:"default_forum_layout,omitempty"`
	Position             int                         `json:"position,omitempty"`
	Reason               string
}

func (c Channel) Edit(Options EditChannelOptions) (*Channel, error) {
	req_body, err := json.Marshal(Options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/channels/%s", API_URL, c.ID), bytes.NewReader(req_body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("X-Audit-Log-Reason", Options.Reason)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	res_body, err := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(res_body))
	}
	var channel Channel
	if err = json.Unmarshal(res_body, &channel); err != nil {
		return nil, err
	}
	return &channel, nil
}
