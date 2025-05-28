package enums

import "godiscord.foo.ng/lib/internal/types"

var InteractionResponseType = struct {
	Ping                           types.InteractionResponseType
	ApplicationCommand             types.InteractionResponseType
	MessageComponent               types.InteractionResponseType
	ApplicationCommandAutocomplete types.InteractionResponseType
	ModalSubmit                    types.InteractionResponseType
}{
	Ping:                           types.IRT_Ping,
	ApplicationCommand:             types.IRT_ApplicationCommand,
	MessageComponent:               types.IRT_MessageComponent,
	ApplicationCommandAutocomplete: types.IRT_ApplicationCommandAutocomplete,
	ModalSubmit:                    types.IRT_ModalSubmit,
}
