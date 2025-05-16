package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/AYn0nyme/godiscord/internal/enums"
)

type Guild struct {
	ID                          string                           `json:"id"`
	Name                        string                           `json:"name"`
	IconHash                    string                           `json:"icon"`
	SplashHash                  string                           `json:"splash"`
	DiscoverySplashHash         string                           `json:"discovery_splash"`
	AmIOwner                    bool                             `json:"owner"`
	OwnerID                     string                           `json:"owner_id"`
	AFKChannelID                string                           `json:"afk_channel_id"`
	AFKTimeout                  int                              `json:"afk_timeout"`
	WidgetEnabled               bool                             `json:"widget_enabled"`
	WidgetChannelID             string                           `json:"wiget_channel_id"`
	VerificationLevel           int                              `json:"verification_level"`
	DefaultMessageNotifications enums.MessageNotificationLevel   `json:"default_message_notifications"`
	ExplicitContentFilter       enums.ExplicitContentFilterLevel `json:"explicit_content_filter"`
	Features                    []enums.GuildFeature             `json:"features"`
	CustomURL                   string                           `json:"vanity_url_code"`
	Description                 string                           `json:"description"`
	BannerHash                  string                           `json:"banner_hash"`
	BoostLevel                  int                              `json:"premium_tier"`
	BoostCount                  int                              `json:"premium_subscription_count"`
	PreferredLocale             string                           `json:"preferred_locale"`
	ApproximateMemberCount      int                              `json:"approximate_member_count"`
	NSFWLevel                   int                              `json:"nsfw_level"`
	BoostProgressionBarEnabled  bool                             `json:"premium_progress_bar_enabled"`
	Roles                       []Role                           `json:"roles"`
	Emojis                      []Emoji                          `json:"emojis"`
	Channels                    []BaseChannel                    `json:"channels"`
	Owner                       GuildMember
	MemberCache                 map[string]GuildMember
	Invites                     InviteManager
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
type BanOptions struct {
	DeleteMessageSeconds int `json:"delete_message_seconds,omitempty"`
	Reason               string
}

// TODO: "refactor" to do good error handling with reao
func (g Guild) Ban(USER any, Options BanOptions) error {
	switch user := USER.(type) {
	case string:
		body, err := json.Marshal(Options)
		if err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/guilds/%s/bans/%s", API_URL, g.ID, user), bytes.NewReader(body))
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
		req.Header.Set("X-Audit-Log-Reason", Options.Reason)
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
		body, err := json.Marshal(Options)
		if err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/guilds/%s/bans/%s", API_URL, g.ID, user.ID), bytes.NewReader(body))
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
		req.Header.Set("X-Audit-Log-Reason", Options.Reason)
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
func (g Guild) BulkBan(Users []string, Options BanOptions) error {
	body, err := json.Marshal(Options)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/guilds/%s/bulk-ban", API_URL, g.ID), bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("X-Audit-Log-Reason", Options.Reason)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		defer res.Body.Close()
		body_in_bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body_in_bytes))
	}
	return nil
}
func (g Guild) UnBan(USER any, Reason ...string) error {
	var reason string
	if len(Reason) > 0 {
		reason = Reason[0]
	}
	switch user := USER.(type) {
	case string:
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/guilds/%s/bans/%s", API_URL, g.ID, user), nil)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
		req.Header.Set("X-Audit-Log-Reason", reason)
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
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/guilds/%s/bans/%s", API_URL, g.ID, user.ID), nil)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
		req.Header.Set("X-Audit-Log-Reason", reason)
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
func (g Guild) Delete() error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/guilds/%s", API_URL, g.ID), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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
func (g Guild) GetMemberByID(ID string) (*GuildMember, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/guilds/%s/members/%s", API_URL, g.ID, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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
func (g Guild) CreateChannel(Options CreateChannelOptions) (*BaseChannel, error) {
	body, err := json.Marshal(Options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/guilds/%s/channels", API_URL, g.ID), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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
func (g Guild) GetBans() (*[]Ban, error) {
	var bans []Ban
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/guilds/%s/bans", API_URL, g.ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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
func (g Guild) GetBan(UserID string) (*Ban, error) {
	var ban Ban
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/guilds/%s/bans/%s", API_URL, g.ID, UserID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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
func (g Guild) CreateRole(Options CreateRoleOptions) (*Role, error) {
	req_body, err := json.Marshal(Options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/guilds/%s/roles", API_URL, g.ID), bytes.NewReader(req_body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("%s", string(body))
	}
	var role Role
	if err = json.Unmarshal(body, &role); err != nil {
		return nil, err
	}
	return &role, nil
}
func (g Guild) DeleteRole(role any, Reason ...string) error {
	var reason string
	if len(Reason) > 0 {
		reason = Reason[0]
	}
	switch ROLE := role.(type) {
	case string:
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/guilds/%s/roles/%s", API_URL, g.ID, ROLE), nil)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
		req.Header.Set("X-Audit-Log-Reason", reason)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusNoContent {
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}
			return fmt.Errorf("%s", string(body))
		}
	case Role:
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/guilds/%s/roles/%s", API_URL, g.ID, ROLE.ID), nil)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
		req.Header.Set("X-Audit-Log-Reason", reason)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusNoContent {
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}
			return fmt.Errorf("%s", string(body))
		}
	default:
		return fmt.Errorf("error: wrong type inserted in role")
	}
	return nil
}
func (g Guild) GetRoles() (*[]Role, error) {
	var roles []Role
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/guilds/%s/roles", API_URL, g.ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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
	err = json.Unmarshal(body, &roles)
	if err != nil {
		return nil, err
	}
	return &roles, nil
}
func (g Guild) GetRole(RoleID string) (*Role, error) {
	var role Role
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/guilds/%s/roles/%s", API_URL, g.ID, RoleID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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
	err = json.Unmarshal(body, &role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (g Guild) EditRole(RoleID string, Options EditRoleOptions, Reason ...string) (*Role, error) {
	var reason string
	if len(Reason) > 0 {
		reason = Reason[0]
	}
	req_body, err := json.Marshal(Options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/guilds/%s/roles/%s", API_URL, g.ID, RoleID), bytes.NewReader(req_body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("X-Audit-Log-Reason", reason)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	res_body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(res_body))
	}
	var role Role
	if err = json.Unmarshal(res_body, &role); err != nil {
		return nil, err
	}
	return &role, nil
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
