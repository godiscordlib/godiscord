package enums

import "godiscord.foo.ng/lib/internal/types"

var InviteType = struct {
	Guild   types.InviteType
	GroupDM types.InviteType
	Friend  types.InviteType
}{
	Guild:   types.IT_Guild,
	GroupDM: types.IT_GroupDM,
	Friend:  types.IT_Friend,
}
