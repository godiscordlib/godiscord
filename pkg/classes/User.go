package classes

type User struct {
	Base
	Username      string
	Discriminator string
	Bot           bool
	Global_Name   *string
	AvatarHash    *string
	BannerHash    *string
	AccentColor   *int
	Locale        *string
	Flags         *int
	VerifiedBot   *bool
}
