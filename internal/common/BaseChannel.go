package common

import "github.com/AYn0nyme/godiscord/internal/enums"

type BaseChannel struct {
	Base
	Type     enums.ChannelType
	GuildID  *string
	Position *int
	// TODO:
	// Add PermissionOverwrites
	Name  *string
	Flags *int
}
