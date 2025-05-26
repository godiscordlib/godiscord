package enums

import "godiscord.foo.ng/lib/internal/types"

var ExplicitContentFilterLevel = struct {
	Disabled            types.ExplicitContentFilterLevel
	MembersWithoutRoles types.ExplicitContentFilterLevel
	AllMembers          types.ExplicitContentFilterLevel
}{
	Disabled:            types.ECFL_Disabled,
	MembersWithoutRoles: types.ECFL_MembersWithoutRoles,
	AllMembers:          types.ECFL_AllMembers,
}
