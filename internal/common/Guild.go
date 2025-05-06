package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AYn0nyme/godiscord/internal/enums"
)

type Guild struct {
	ID                          string               `json:"id"`
	Name                        string               `json:"name"`
	IconHash                    string               `json:"icon"`
	SplashHash                  string               `json:"splash"`
	DiscoverySplashHash         string               `json:"discovery_splash"`
	AmIOwner                    bool                 `json:"owner"`
	OwnerID                     string               `json:"owner_id"`
	AFKChannelID                string               `json:"afk_channel_id"`
	AFKTimeout                  int                  `json:"afk_timeout"`
	WidgetEnabled               bool                 `json:"widget_enabled"`
	WidgetChannelID             string               `json:"wiget_channel_id"`
	VerificationLevel           int                  `json:"verification_level"`
	DefaultMessageNotifications int                  `json:"default_message_notifications"`
	ExplicitContentFilter       int                  `json:"explicit_content_filter"`
	Features                    []enums.GuildFeature `json:"features"`
	CustomURL                   string               `json:"vanity_url_code"`
	Description                 string               `json:"description"`
	BannerHash                  string               `json:"banner_hash"`
	BoostLevel                  int                  `json:"premium_tier"`
	BoostCount                  int                  `json:"premium_subscription_count"`
	PreferredLocale             string               `json:"preferred_locale"`
	ApproximateMemberCount      int                  `json:"approximate_member_count"`
	NSFWLevel                   int                  `json:"nsfw_level"`
	BoostProgressionBarEnabled  bool                 `json:"premium_progress_bar_enabled"`
	Roles                       []Role               `json:"roles"`
	Emojis                      []Emoji              `json:"emojis"`
	Channels                    []BaseChannel        `json:"channels"`
	Owner                       GuildMember
	MemberCache                 map[string]GuildMember
	// Has2FARequired              bool                 `json:"channels"`
	// TODO:
	// Add Owner
	// Add Members? Will call API several times for a big server (1k members per request)
}

type CreateGuildOptions struct {
	Name                       string                           `json:"name"`
	Icon                       string                           `json:"icon,omitempty"`
	VerificationLevel          enums.VerificationLevel          `json:"verification_level,omitempty"`
	DefaultMessageNotification enums.MessageNotificationLevel   `json:"default_message_notifications,omitempty"`
	ExplicitLevelFilter        enums.ExplicitContentFilterLevel `json:"explicit_content_filter,omitempty"`
	Roles                      []Role                           `json:"roles,omitempty"`
	Channels                   []Channel                        `json:"channels,omitempty"`
	AFKChannelID               string                           `json:"afk_channel_id,omitempty"`
	AFKTimeout                 int                              `json:"afk_timeout,omitempty"`
	SystemChannelID            string                           `json:"system_channel_id,omitempty"`
	SystemChannelFlags         enums.SystemChannelFlag          `json:"system_channel_flags,omitempty"`
}
type EditGuildOptions struct {
	AFKChannelID               string                           `json:"afk_channel_id,omitempty"`
	AFKTimeout                 int                              `json:"afk_timeout,omitempty"`
	ExplicitLevelFilter        enums.ExplicitContentFilterLevel `json:"explicit_content_filter,omitempty"`
	Name                       string                           `json:"name"`
	Icon                       string                           `json:"icon,omitempty"`
	VerificationLevel          enums.VerificationLevel          `json:"verification_level,omitempty"`
	DefaultMessageNotification enums.MessageNotificationLevel   `json:"default_message_notifications,omitempty"`
	Roles                      []Role                           `json:"roles,omitempty"`
	Channels                   []Channel                        `json:"channels,omitempty"`
	SystemChannelID            string                           `json:"system_channel_id,omitempty"`
	SystemChannelFlags         enums.SystemChannelFlag          `json:"system_channel_flags,omitempty"`
	OwnerID                    string                           `json:"owner_id,omitempty"` // ONLY IF THE BOT IS THE OWNER OF THE GUILD
	Splash                     string                           `json:"splash"`
	DiscoverySplash            string                           `json:"discovery_splash,omitempty"`
	Banner                     string                           `json:"banner,omitempty"`
	RulesChannelID             string                           `json:"rules_channel_id,omitempty"`          // Only on community guilds
	PublicUpdatesChannelID     string                           `json:"public_updates_channel_id,omitempty"` // Only on community guilds
	PreferredLocale            string                           `json:"preferred_locale,omitempty"`          // Only on community guilds
	Description                string                           `json:"description,omitempty"`               // Only on community guilds
	SafetyAlertsChannelID      string                           `json:"safety_alerts_channel_id,omitempty"`
	Features                   []enums.GuildFeature             `json:"features"`
	BoostProgressionBarEnabled bool                             `json:"premium_progress_bar_enabled,omitempty"`
}

func (g Guild) Ban(Client Client, USER any, DeleteMessageSeconds int, Reason string) error {
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
		req.Header.Set("X-Audit-Log-Reason", Reason)
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
		req.Header.Set("X-Audit-Log-Reason", Reason)
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
func (g Guild) UnBan(Client Client, USER any, DeleteMessageSeconds int, Reason string) error {
	switch user := USER.(type) {
	case string:
		var body bytes.Buffer
		json.NewEncoder(&body).Encode(struct {
			DMS int `json:"delete_message_seconds"`
		}{
			DMS: DeleteMessageSeconds,
		})
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/guilds/%s/bans/%s", API_URL, g.ID, user), &body)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bot "+Client.Token)
		req.Header.Set("X-Audit-Log-Reason", Reason)
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

		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/guilds/%s/bans/%s", API_URL, g.ID, user.ID), &body)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bot "+Client.Token)
		req.Header.Set("X-Audit-Log-Reason", Reason)
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
func (g Guild) Delete(Client Client) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/guilds/%s", API_URL, g.ID), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bot "+Client.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 204 {
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("%s", string(body))
	}
	return nil
}
func (g Guild) GetMemberByID(Client Client, ID string) (*GuildMember, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/guilds/%s/members/%s", API_URL, g.ID, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+Client.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s", string(body))
	}
	var gm GuildMember
	json.Unmarshal(body, &gm)
	return &gm, nil
}
func (g Guild) CreateChannel(Client Client, Options CreateChannelOptions) (*BaseChannel, error) {
	body, err := json.Marshal(Options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/guilds/%s/channels", API_URL, g.ID), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+Client.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	req_body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("error: %s", string(req_body))
	}
	var channel BaseChannel
	err = json.Unmarshal(req_body, &channel)
	if err != nil {
		return nil, err
	}
	return &channel, nil
}
func (g Guild) GetBans(Client Client) (*[]Ban, error) {
	var bans []Ban
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/guilds/%s/bans", API_URL, g.ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+Client.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", string(body))
	}
	err = json.Unmarshal(body, &bans)
	if err != nil {
		return nil, err
	}
	return &bans, nil
}
func (g Guild) GetBan(Client Client, UserID string) (*Ban, error) {
	var ban Ban
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/guilds/%s/bans/%s", API_URL, g.ID, UserID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+Client.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", string(body))
	}
	err = json.Unmarshal(body, &ban)
	if err != nil {
		return nil, err
	}
	return &ban, nil
}

// WIP
// func (g Guild) Edit(Client Client, Options EditGuildOptions) (*Guild, error) {
// 	body, err := json.Marshal(Options)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/guilds/%s", API_URL, g.ID), bytes.NewReader(body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Authorization", "Bot "+Client.Token)
// 	req.Header.Set("Content-Type", "application/json")
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if res.StatusCode != http.StatusOK {
// 		return nil, err
// 	}
// 	defer res.Body.Close()
// 	body_in_bytes, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var guild Guild
// 	json.Unmarshal(body_in_bytes, &guild)
// 	return &guild, nil
// }
