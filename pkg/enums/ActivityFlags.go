package enums

import (
	"godiscord.foo.ng/lib/pkg/types"
)

var ActivityFlag = struct {
	Instance    types.ActivityFlag
	Join        types.ActivityFlag
	Spectate    types.ActivityFlag
	JoinRequest types.ActivityFlag
	Sync        types.ActivityFlag
}{
	Instance:    types.AF_Instance,
	Join:        types.AF_Join,
	Spectate:    types.AF_Spectate,
	JoinRequest: types.AF_JoinRequest,
	Sync:        types.AF_Sync,
}
