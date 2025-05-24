package enums

type ActivityType int

const (
	AT_Playing   ActivityType = iota // 0
	AT_Streaming                     // 1
	AT_Listening                     // 2
	AT_Watching                      // 3
	AT_Custom                        // 4 (par exemple)
)
