package classes

import (
	types2 "godiscord.foo.ng/lib/pkg/types"
)

type BaseInteraction struct {
	Type   types2.InteractionResponseType `json:"type"`
	Token  string                         `json:"token"`
	Member GuildMember                    `json:"member"`
	ID     string                         `json:"id"`
	Guild  Guild                          `json:"guild"`
	Data   baseInteractionData            `json:"data"`
}
type baseInteractionData struct {
	Type types2.InteractionType `json:"type"`
	Name string                 `json:"name"`
	ID   string                 `json:"id"`
}

type BaseComponent interface {
	GetType() types2.InteractionType
}
