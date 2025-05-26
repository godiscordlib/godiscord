package types

type ActivityType int

const (
	AT_Playing ActivityType = iota
	AT_Streaming
	AT_Listening
	AT_Watching
	AT_Custom
)
