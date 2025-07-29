package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var ExplicitContentFilterLevel = struct {
	Disabled            types.ExplicitContentFilterLevel
	MembersWithoutRoles types.ExplicitContentFilterLevel
	AllMembers          types.ExplicitContentFilterLevel
}{
	Disabled:            types.ECFL_Disabled,
	MembersWithoutRoles: types.ECFL_MembersWithoutRoles,
	AllMembers:          types.ECFL_AllMembers,
}
