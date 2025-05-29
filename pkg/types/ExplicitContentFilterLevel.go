package types

type ExplicitContentFilterLevel int

const (
	ECFL_Disabled ExplicitContentFilterLevel = iota
	ECFL_MembersWithoutRoles
	ECFL_AllMembers
)
