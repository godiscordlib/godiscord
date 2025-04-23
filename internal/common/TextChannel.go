package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TextChannel struct {
	BaseChannel
	LastMessageID string
	NSFW          bool
	CategoryID    string
}

func (t TextChannel) Send(Client Client, Data any) {
	switch data := Data.(type) {
	case string:
		message := payloadMessage{
			Content: data,
		}

		var payload bytes.Buffer
		json.NewEncoder(&payload).Encode(message)

		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages", API_URL, t.ID), &payload)
		request.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		request.Header.Add("Content-Type", "application/json")
		fmt.Println(request)
		http.DefaultClient.Do(request)
	case MessageData:
		message := payloadMessage{
			Content: data.Content,
			Embeds:  data.Embeds,
		}

		var payload bytes.Buffer
		json.NewEncoder(&payload).Encode(message)
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages", API_URL, t.ID), &payload)
		request.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		request.Header.Add("Content-Type", "application/json")
		http.DefaultClient.Do(request)
	}
}
