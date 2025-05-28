package classes

import "godiscord.foo.ng/lib/internal/types"

type BaseInteraction struct {
	Type   types.InteractionType `json:"type"`
	Token  string                `json:"token"`
	Member GuildMember           `json:"member"`
	ID     string                `json:"id"`
}

type BaseComponent interface {
	GetType() types.InteractionType
}
