package classes

import "strconv"

type Embed struct {
	Title       string         `json:"title"`
	Type        string         `json:"type"`
	Description string         `json:"description"`
	URL         string         `json:"url"`
	Thumbnail   EmbedThumbnail `json:"thumbnail"`
	Image       EmbedImage     `json:"image"`
	Timestamp   string         `json:"timestamp"` // ISO8601 timestamp
	Color       int            `json:"color"`
	Footer      EmbedFooter    `json:"footer"`
	Fields      []EmbedField   `json:"fields"`
}
type EmbedThumbnail struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
type EmbedImage struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
type EmbedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
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

// Set the color of the embed from Hex Code. Don't include the #
func (e Embed) SetColor(HexColor string) Embed {
	color_in_int, err := strconv.ParseInt(HexColor, 16, 64)
	if err != nil {
		panic(err)
	}
	e.Color = int(color_in_int)
	return e
}
func (e Embed) SetFooter(FooterConfig EmbedFooter) Embed {
	e.Footer.IconURL = FooterConfig.IconURL
	e.Footer.Text = FooterConfig.Text
	return e
}
func (e Embed) AddField(Name, Value string, Inline bool) Embed {
	e.Fields = append(e.Fields, EmbedField{Name: Name, Value: Value, Inline: Inline})
	return e
}
func (e Embed) SetThumbnail(ThumbnailConfig EmbedThumbnail) Embed {
	e.Thumbnail.URL = ThumbnailConfig.URL
	e.Thumbnail.Height = ThumbnailConfig.Height
	e.Thumbnail.Width = ThumbnailConfig.Width
	return e
}

func (e Embed) SetImage(ImageConfig EmbedImage) Embed {
	e.Image.URL = ImageConfig.URL
	e.Image.Height = ImageConfig.Height
	e.Image.Width = ImageConfig.Width
	return e
}
