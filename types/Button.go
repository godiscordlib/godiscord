package types

type Button struct {
	Type     int
	Style    int
	Label    *string
	Emoji    *string
	CustomID *string
	SKUID    *string
	URL      *string
	Disabled *bool
}
