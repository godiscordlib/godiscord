package classes

import (
	"godiscord.foo.ng/lib/pkg/types"
)

type Invite struct {
	Type      types.InviteType `json:"type"`
	Code      string           `json:"code"`
	Inviter   User             `json:"inviter"`
	ExpiresAt string           `json:"expires_at"` // ISO8601 timestamp
}
