package enums

import (
	"godiscord.foo.ng/lib/pkg/types"
)

var SystemChannelFlag = struct {
	SuppressJoinNotifications                           types.SystemChannelFlag
	SuppressPremiumSubscriptions                        types.SystemChannelFlag
	SuppressGuildReminderNotifications                  types.SystemChannelFlag
	SuppressJoinNotificationReplies                     types.SystemChannelFlag
	SuppressRoleSubscriptionPurchaseNotifications       types.SystemChannelFlag
	SuppressRoleSubscriptionPurchaseNotificationReplies types.SystemChannelFlag
}{
	SuppressJoinNotifications:                           types.SCF_SuppressJoinNotifications,
	SuppressPremiumSubscriptions:                        types.SCF_SuppressPremiumSubscriptions,
	SuppressGuildReminderNotifications:                  types.SCF_SuppressGuildReminderNotifications,
	SuppressJoinNotificationReplies:                     types.SCF_SuppressJoinNotificationReplies,
	SuppressRoleSubscriptionPurchaseNotifications:       types.SCF_SuppressRoleSubscriptionPurchaseNotifications,
	SuppressRoleSubscriptionPurchaseNotificationReplies: types.SCF_SuppressRoleSubscriptionPurchaseNotificationReplies,
}
