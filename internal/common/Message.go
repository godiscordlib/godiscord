package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Base
	ChannelID        string   `json:"channel_id"`
	Author           User     `json:"author"`
	Content          string   `json:"content"`
	Timestamp        string   `json:"timestamp"` // ISO8601 timestamp
	MentionsEveryone bool     `json:"mention_everyone"`
	UsersMentions    []User   `json:"mentions"`
	RolesMentions    []string `json:"mention_roles"`
	Embeds           []Embed  `json:"embeds"`
	Pinned           bool     `json:"pinned"`
	Type             int      `json:"type"`
	Flags            int      `json:"flags"`
}
type MessageData struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
	Flags   int     `json:"flags"`
}
type payloadMessage struct {
	Content   string                  `json:"content"`
	Embeds    []Embed                 `json:"embeds"`
	Reference payloadMessageReference `json:"message_reference"`
}
type payloadMessageReference struct {
	ID   string `json:"message_id"`
	Type int    `json:"type"`
}

const API_URL = "https://discord.com/api/v10"

func (m Message) Reply(Client Client, Data any) {
	switch data := Data.(type) {
	case string:
		message := payloadMessage{
			Content: data,
			Reference: payloadMessageReference{
				ID:   m.ID,
				Type: 0,
			},
		}

		var payload bytes.Buffer
		json.NewEncoder(&payload).Encode(message)

		httpClient := http.Client{}
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages", API_URL, m.ChannelID), &payload)
		request.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		request.Header.Add("Content-Type", "application/json")
		httpClient.Do(request)
	case MessageData:
		message := payloadMessage{
			Content: data.Content,
			Embeds:  data.Embeds,
			Reference: payloadMessageReference{
				ID:   m.ID,
				Type: 0,
			},
		}

		var payload bytes.Buffer
		json.NewEncoder(&payload).Encode(message)
		httpClient := http.Client{}
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages", API_URL, m.ChannelID), &payload)
		request.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		request.Header.Add("Content-Type", "application/json")
		httpClient.Do(request)
	}
}
