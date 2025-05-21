package common

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GuildMember struct {
	User          User     `json:"user"`
	Nickname      string   `json:"nick"`
	AvatarHash    string   `json:"avatar"`
	BannerHash    string   `json:"banner"`
	Roles         []string `json:"string"`
	JoinedAt      string   `json:"joined_at"`     //	ISO8601 timestamp
	BoostingSince string   `json:"premium_since"` // ISO8601 timestamp
	Deafened      bool     `json:"deaf"`
	Mute          bool     `json:"mute"`
	StillJoining  bool     `json:"pending"`                      // If the user is still on the Membership screening
	Permissions   string   `json:"permissions"`                  // TODO: check if it's possible to juste replace with Permissions from the internal/enums.
	TimedoutUntil string   `json:"communication_disabled_until"` // ISO8601 timestamp
	RoleManager   guildMemberRoleManager
	// TODO:
	// Add GuildMemberFlags
	// Add AvatarDecoration (low priority)
}

type guildMemberRoleManager struct {
	MemberID string
	GuildID  string
}

func (rm guildMemberRoleManager) Add(RoleID string, Reason ...string) error {
	var reason string
	if len(Reason) > 1 {
		reason = Reason[0]
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/guilds/%s/members/%s/roles/%s", API_URL, rm.GuildID, rm.MemberID, RoleID), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("X-Audit-Log-Reason", reason)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 204 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	return nil
}

type EditGuildMemberOptions struct {
	Nickname        string   `json:"nick"`
	Roles           []string `json:"roles"`
	Muted           bool     `json:"muted"`
	Deafened        bool     `json:"deafened"`
	MoveToChannelID string   `json:"channel_id"`
	TimeoutUntil    string   `json:"communication_disabled_until"` // ISO8601 timestamp
	Flags           int      `json:"flags"`
}

func (gm GuildMember) Edit(Options EditGuildMemberOptions) {

}
