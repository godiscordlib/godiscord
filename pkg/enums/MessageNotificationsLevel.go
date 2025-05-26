package enums

import "godiscord.foo.ng/lib/internal/types"

var MessageNotificationLevel = struct {
	AllMessages  types.MessageNotificationLevel
	OnlyMentions types.MessageNotificationLevel
}{
	AllMessages:  types.MNL_AllMessages,
	OnlyMentions: types.MNL_OnlyMentions,
}
