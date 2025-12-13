package classes

type User struct {
	Base
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Bot           bool   `json:"bot"`
	Global_Name   string `json:"display_username"`
	AvatarHash    string `json:"avatar"`
	BannerHash    string `json:"banner"`
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
