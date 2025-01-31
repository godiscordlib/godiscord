package types

type User struct {
	Base
	Username      string
	Discriminator string
	Global_Name   *string
	AvatarHash    *string
	Bot           *bool
	BannerHash    *string
	AccentColor   *int
	Locale        *string
	Flags         *int
}
