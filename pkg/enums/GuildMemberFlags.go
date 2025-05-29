package enums

import (
	"godiscord.foo.ng/lib/pkg/types"
)

var GuildMemberFlag = struct {
	DidRejoin                    types.GuildMemberFlag
	CompletedOnBoarding          types.GuildMemberFlag
	BypassesVerification         types.GuildMemberFlag
	StartedOnBoarding            types.GuildMemberFlag
	IsGuest                      types.GuildMemberFlag
	StartedHomeActions           types.GuildMemberFlag
	CompletedHomeActions         types.GuildMemberFlag
	AutoModQuarantinedUsername   types.GuildMemberFlag
	DMSettingsUpsellAcknowledged types.GuildMemberFlag
}{
	DidRejoin:                    types.GMF_DidRejoin,
	CompletedOnBoarding:          types.GMF_CompletedOnBoarding,
	BypassesVerification:         types.GMF_BypassesVerification,
	StartedOnBoarding:            types.GMF_StartedOnBoarding,
	IsGuest:                      types.GMF_IsGuest,
	StartedHomeActions:           types.GMF_StartedHomeActions,
	CompletedHomeActions:         types.GMF_CompletedHomeActions,
	AutoModQuarantinedUsername:   types.GMF_AutoModQuarantinedUsername,
	DMSettingsUpsellAcknowledged: types.GMF_DMSettingsUpsellAcknowledged,
}
