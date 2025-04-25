package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AYn0nyme/godiscord/internal/enums"
)

type Guild struct {
	ID                          string          `json:"id"`
	Name                        string          `json:"name"`
	IconHash                    string          `json:"icon"`
	SplashHash                  string          `json:"splash"`
	DiscoverySplashHash         string          `json:"discovery_splash"`
	AmIOwner                    bool            `json:"owner"`
	OwnerID                     string          `json:"owner_id"`
	AFKChannelId                string          `json:"afk_channel_id"`
	AFKTimeout                  int             `json:"afk_timeout"`
	WidgetEnabled               bool            `json:"widget_enabled"`
	WidgetChannelID             string          `json:"wiget_channel_id"`
	VerificationLevel           int             `json:"verification_level"`
	DefaultMessageNotifications int             `json:"default_message_notifications"`
	ExplicitContentFilter       int             `json:"explicit_content_filter"`
	Features                    []enums.Feature `json:"features"`
	CustomURL                   string          `json:"vanity_url_code"`
	Description                 string          `json:"description"`
	BannerHash                  string          `json:"banner_hash"`
	BoostLevel                  int             `json:"premium_tier"`
	BoostCount                  int             `json:"premium_subscription_count"`
	PreferredLocale             string          `json:"preferred_locale"`
	ApproximateMemberCount      int             `json:"approximate_member_count"`
	NSFWLevel                   int             `json:"nsfw_level"`
	BoostProgressionBarEnabled  bool            `json:"premium_progress_bar_enabled"`
	Roles                       []Role          `json:"roles"`
	Emojis                      []Emoji         `json:"emojis"`
	Channels                    []BaseChannel
	Has2FARequired              bool
	// TODO:
	// Add Owner
	// Add Members? Will call API several times for a big server (1k members per request)
}

func (g Guild) Ban(Client Client, USER any, DeleteMessageSeconds int) error {
	switch user := USER.(type) {
	case string:
		var body bytes.Buffer
		json.NewEncoder(&body).Encode(struct {
			DMS int `json:"delete_message_seconds"`
		}{
			DMS: DeleteMessageSeconds,
		})
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/guilds/%s/bans/%s", API_URL, g.ID, user), &body)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bot "+Client.Token)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode == 403 {
			return fmt.Errorf("error: not enough permissions to ban user")
		}
		if res.StatusCode != 204 {
			return fmt.Errorf("error: ban resulted in a %s code, instead of 204", res.Status)
		}
	case User:
		var body bytes.Buffer
		json.NewEncoder(&body).Encode(struct {
			DMS int `json:"delete_message_seconds"`
		}{
			DMS: DeleteMessageSeconds,
		})

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/guilds/%s/bans/%s", API_URL, g.ID, user.ID), &body)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bot "+Client.Token)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode == 403 {
			return fmt.Errorf("error: not enough permissions to ban user")
		}
		if res.StatusCode != 204 {
			return fmt.Errorf("error: ban resulted in a %s code, instead of 204", res.Status)
		}
	default:
		return fmt.Errorf("error: User is not a string (ID) nor a User struct")
	}
	return nil
}
