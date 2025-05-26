package types

type InviteType int

const (
	IT_Guild InviteType = iota
	IT_GroupDM
	IT_Friend
)
