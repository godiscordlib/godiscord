package types

type InteractionResponseType int

const (
	IRT_Ping InteractionResponseType = iota + 1
	IRT_ApplicationCommand
	IRT_MessageComponent
	IRT_ApplicationCommandAutocomplete
	IRT_ModalSubmit
)
