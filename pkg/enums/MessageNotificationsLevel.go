package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var MessageNotificationLevel = struct {
	AllMessages  types.MessageNotificationLevel
	OnlyMentions types.MessageNotificationLevel
}{
	AllMessages:  types.MNL_AllMessages,
	OnlyMentions: types.MNL_OnlyMentions,
}
