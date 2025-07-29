package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var MessageFlags = struct {
	Crossposted                      types.MessageFlag
	IsCrosspost                      types.MessageFlag
	SuppressEmbeds                   types.MessageFlag
	SourceMessageDeleted             types.MessageFlag
	Urgent                           types.MessageFlag
	HasThread                        types.MessageFlag
	Ephemeral                        types.MessageFlag
	Loading                          types.MessageFlag
	FailedToMentionSomeRolesInThread types.MessageFlag
	SuppressNotifications            types.MessageFlag
	IsVoiceMessage                   types.MessageFlag
	HasSnapshot                      types.MessageFlag
	IsComponentsV2                   types.MessageFlag
}{
	types.MF_Crossposted,
	types.MF_IsCrosspost,
	types.MF_SuppressEmbeds,
	types.MF_SourceMessageDeleted,
	types.MF_Urgent,
	types.MF_HasThread,
	types.MF_Ephemeral,
	types.MF_Loading,
	types.MF_FailedToMentionSomeRolesInThread,
	types.MF_SuppressNotifications,
	types.MF_IsVoiceMessage,
	types.MF_HasSnapshot,
	types.MF_IsComponentsV2,
}
