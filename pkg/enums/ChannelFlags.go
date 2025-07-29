package enums

import "github.com/godiscordlib/godiscord/pkg/types"

var ChannelFlag = struct {
	Pinned                   types.ChannelFlag
	RequireTag               types.ChannelFlag
	HideMediaDownloadOptions types.ChannelFlag
}{
	1 << 1,
	1 << 4,
	1 << 15,
}
