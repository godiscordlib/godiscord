package common

type Reaction struct {
	Count             int             `json:"count"`
	CountDetails      ReactionDetails `json:"count_details"`
	HaveIReacted      bool            `json:"me"`
	HaveISuperReacted bool            `json:"me_burst"`
	Emoji             Emoji           `json:"emoji"`
	SuperColors       []int           `json:"burst_colors"`
}
type ReactionDetails struct {
	Super  int `json:"burst"`  // Super reactions
	Normal int `json:"normal"` // Normal reactions
}
