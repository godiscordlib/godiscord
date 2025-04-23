package common

type PermissionOverwrite struct {
	Base
	Type  int    `json:"type"`
	Allow string `json:"allow"`
	Deny  string `json:"deny"`
}
