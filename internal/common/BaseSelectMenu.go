package common

type BaseSelectMenu struct {
	Type        int
	CustomID    string
	Placeholder *string
	MinValues   int
	MaxValues   int
	Disabled    bool
}
