package classes

type User struct {
	Base
	Username      string
	Discriminator string
	Bot           bool
	Global_Name   string
	AvatarHash    string
	BannerHash    string
	AccentColor   int
	Locale        string
	Flags         int
	VerifiedBot   bool
}

/*
Get the avatar URL of the user

Type can either be "png", "jpg" or "gif"; defaults with png
*/
func (u User) GetAvatarURL(Type string) string {
	if Type == "" {
		Type = "png"
	}
	return "https://cdn.discordapp.com/avatars/" + u.ID + "/" + u.AvatarHash + "." + Type
}
