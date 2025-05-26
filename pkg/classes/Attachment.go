package classes

import "godiscord.foo.ng/lib/internal/types"

type Attachment struct {
	ID          int                  `json:"id"`
	FileName    string               `json:"filename"`
	Title       string               `json:"title,omitempty"`
	Description string               `json:"description,omitempty"`
	ContentType string               `json:"content_type,omitempty"`
	Size        int                  `json:"size,omitempty"`
	URL         string               `json:"url,omitempty"`
	ProxyURL    string               `json:"proxy_url,omitempty"`
	Height      int                  `json:"height,omitempty"`
	Width       int                  `json:"width,omitempty"`
	Ephemeral   bool                 `json:"ephemeral,omitempty"`
	Duration    int                  `json:"duration_secs,omitempty"`
	Waveform    string               `json:"waveform,omitempty"`
	Flags       types.AttachmentFlag `json:"flags,omitempty"`
	FilePath    string
}
