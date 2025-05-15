package common

import "github.com/AYn0nyme/godiscord/internal/enums"

type Invite struct {
	Type      enums.InviteType `json:"type"`
	Code      string           `json:"code"`
	Inviter   User             `json:"inviter"`
	ExpiresAt string           `json:"expires_at"` // ISO8601 timestamp
}
