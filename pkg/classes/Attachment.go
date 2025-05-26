package classes

import "godiscord.foo.ng/lib/internal/types"

type Attachment struct {
	FileName    string               `json:"filename"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	ContentType string               `json:"content_type"`
	Size        int                  `json:"size"`
	URL         string               `json:"url"`
	ProxyURL    string               `json:"proxy_url"`
	Height      int                  `json:"height"`
	Width       int                  `json:"width"`
	Ephemeral   bool                 `json:"ephemeral"`
	Duration    int                  `json:"duration_secs"`
	Waveform    string               `json:"waveform"`
	Flags       types.AttachmentFlag `json:"flags"`
}
