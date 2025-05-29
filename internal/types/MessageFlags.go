package types

type MessageFlag int

const (
	MF_Crossposted MessageFlag = 1 << iota
	MF_IsCrosspost
	MF_SuppressEmbeds
	MF_SourceMessageDeleted
	MF_Urgent
	MF_HasThread
	MF_Ephemeral
	MF_Loading
	MF_FailedToMentionSomeRolesInThread
)
const (
	MF_SuppressNotifications MessageFlag = 1 << 12
	MF_IsVoiceMessage        MessageFlag = 1 << 13
	MF_HasSnapshot           MessageFlag = 1 << 14
	MF_IsComponentsV2        MessageFlag = 1 << 15
)
