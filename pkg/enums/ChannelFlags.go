package enums

import "godiscord.foo.ng/lib/pkg/types"

var ChannelFlag = struct {
	Pinned                   types.ChannelFlag
	RequireTag               types.ChannelFlag
	HideMediaDownloadOptions types.ChannelFlag
}{
	1 << 1,
	1 << 4,
	1 << 15,
}
