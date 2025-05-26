package classes

type Ban struct {
	Reason string `json:"reason"`
	User   User   `json:"user"`
}
