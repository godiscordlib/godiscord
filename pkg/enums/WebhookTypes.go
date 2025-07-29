package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var WebhookType = struct {
	Incomming       types.WebhookType
	ChannelFollower types.WebhookType
	Application     types.WebhookType
}{
	Incomming:       types.WT_Incomming,
	ChannelFollower: types.WT_ChannelFollower,
	Application:     types.WT_Application,
}
