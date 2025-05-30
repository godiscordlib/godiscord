package types

type ChannelFlag int

const (
	CF_Pinned                   ChannelFlag = 1 << 1
	CF_RequireTag               ChannelFlag = 1 << 4
	CF_HideMediaDownloadOptions ChannelFlag = 1 << 15
)
