package classes

import "godiscord.foo.ng/lib/internal/types"

type BaseInteraction struct {
	Type   types.InteractionType `json:"type"`
	Token  string                `json:"token"`
	Member GuildMember           `json:"member"`
	ID     string                `json:"id"`
	Guild  Guild                 `json:"guild"`
	Data   baseInteractionData   `json:"data"`
}
type baseInteractionData struct {
	Type types.InteractionType `json:"type"`
	Name string                `json:"name"`
	ID   string                `json:"id"`
}

type BaseComponent interface {
	GetType() types.InteractionType
}
