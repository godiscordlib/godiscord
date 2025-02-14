package types

import "godiscord/enums"

type Guild struct {
	Base
	Name                        string
	IconHash                    *string
	SplashHash                  *string
	DiscoverySplashHash         *string
	AmIOwner                    *bool
	OwnerID                     string
	AFKChannelId                *string
	AFKTimeout                  int
	WidgetEnabled               *bool
	WidgetChannelID             string
	VerificationLevel           int
	DefaultMessageNotifications int
	ExplicitContentFilter       int
	Features                    []enums.Feature
	Has2FARequired              bool
	CustomURL                   *string
	Description                 *string
	BannerHash                  *string
	BoostLevel                  int
	BoostCount                  int
	PreferredLocale             *string
	ApproximateMemberCount      *int
	NSFWLevel                   int
	BoostProgressionBarEnabled  bool
	Roles                       []Role
	Emojis                      []Emoji
}
