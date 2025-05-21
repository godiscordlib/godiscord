package enums

type GuildMemberFlag int

const (
	GMF_DidRejoin GuildMemberFlag = 1 << iota
	GMF_CompletedOnBoarding
	GMF_BypassesVerification
	GMF_StartedOnBoarding
	GMF_IsGuest
	GMF_StartedHomeActions
	GMF_CompletedHomeActions
	GMF_AutoModQuarantinedUsername
	GMF_DMSettingsUpsellAcknowledged
)
