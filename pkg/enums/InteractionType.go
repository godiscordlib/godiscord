package enums

import (
	"godiscord.foo.ng/lib/pkg/types"
)

var InteractionType = struct {
	ChatInput         types.InteractionType
	User              types.InteractionType
	Message           types.InteractionType
	PrimaryEntryPoint types.InteractionType
}{
	ChatInput:         types.IT_ChatInput,
	User:              types.IT_User,
	Message:           types.IT_Message,
	PrimaryEntryPoint: types.IT_PrimaryEntryPoint,
}
