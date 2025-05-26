package enums

import "godiscord.foo.ng/lib/internal/types"

var WebhookType = struct {
	Incomming       types.WebhookType
	ChannelFollower types.WebhookType
	Application     types.WebhookType
}{
	Incomming:       types.WT_Incomming,
	ChannelFollower: types.WT_ChannelFollower,
	Application:     types.WT_Application,
}
