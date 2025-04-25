package common

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
	// TODO:
	// Add GuildMemberFlags
	// Add AvatarDecoration (low priority)
}
