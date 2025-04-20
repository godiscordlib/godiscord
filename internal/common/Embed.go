package common

type Embed struct {
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Thumbnail   thumnbail `json:"thumbnail"`
	Timestamp   string    `json:"timestamp"` // ISO8601 timestamp
	Color       int       `json:"color"`
	Footer      footer    `json:"footer"`
	Fields      []field   `json:"fields"`
}
type thumnbail struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
type footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

func NewEmbed() Embed {
	return Embed{Color: 2895667}
}

func (e Embed) SetTitle(Title string) Embed {
	e.Title = Title
	return e
}
func (e Embed) SetDescription(Description string) Embed {
	e.Description = Description
	return e
}
func (e Embed) SetURL(URL string) Embed {
	e.URL = URL
	return e
}
func (e Embed) SetTimestamp(Timestamp string) Embed {
	e.Timestamp = Timestamp
	return e
}
func (e Embed) SetColor(Color int) Embed {
	e.Color = Color
	return e
}
func (e Embed) SetFooter(Text, IconURL string) Embed {
	e.Footer.IconURL = IconURL
	e.Footer.Text = Text
	return e
}
func (e Embed) AddField(Name, Value string, Inline bool) Embed {
	e.Fields = append(e.Fields, field{Name: Name, Value: Value, Inline: Inline})
	return e
}
func (e Embed) SetThumbnail(URL string, Height, Width int) Embed {
	e.Thumbnail.URL = URL
	e.Thumbnail.Height = Height
	e.Thumbnail.Width = Width
	return e
}
