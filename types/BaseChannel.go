package types

type BaseChannel struct {
	Base
	Type     int
	GuildID  *string
	Position *int
	// TODO:
	// Add PermissionOverwrites
	Name  *string
	Flags *int
}
