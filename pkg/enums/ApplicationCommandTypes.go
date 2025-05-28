package enums

import "godiscord.foo.ng/lib/internal/types"

var ApplicationCommandType = struct {
	ChatInput         types.ApplicationCommandType
	User              types.ApplicationCommandType
	Message           types.ApplicationCommandType
	PrimaryEntryPoint types.ApplicationCommandType
}{
	ChatInput:         types.ACT_ChatInput,
	User:              types.ACT_User,
	Message:           types.ACT_Message,
	PrimaryEntryPoint: types.ACT_PrimaryEntryPoint,
}
