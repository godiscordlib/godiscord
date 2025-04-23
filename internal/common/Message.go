package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

var emojiRanges = []unicode.RangeTable{
	{ // Emoticons (U+1F600 - U+1F64F)
		R32: []unicode.Range32{
			{Lo: 0x1F600, Hi: 0x1F64F, Stride: 1},
		},
	},
	{ // Misc Symbols and Pictographs (U+1F300 - U+1F5FF)
		R32: []unicode.Range32{
			{Lo: 0x1F300, Hi: 0x1F5FF, Stride: 1},
		},
	},
	{ // Transport and Map Symbols (U+1F680 - U+1F6FF)
		R32: []unicode.Range32{
			{Lo: 0x1F680, Hi: 0x1F6FF, Stride: 1},
		},
	},
	{ // Supplemental Symbols and Pictographs (U+1F900 - U+1F9FF)
		R32: []unicode.Range32{
			{Lo: 0x1F900, Hi: 0x1F9FF, Stride: 1},
		},
	},
	{ // Symbols and Pictographs Extended-A (U+1FA70 - U+1FAFF)
		R32: []unicode.Range32{
			{Lo: 0x1FA70, Hi: 0x1FAFF, Stride: 1},
		},
	},
	{ // Dingbats (U+2700 - U+27BF)
		R16: []unicode.Range16{
			{Lo: 0x2700, Hi: 0x27BF, Stride: 1},
		},
	},
	{ // Misc Symbols (U+2600 - U+26FF)
		R16: []unicode.Range16{
			{Lo: 0x2600, Hi: 0x26FF, Stride: 1},
		},
	},
}

type Message struct {
	Base
	Channel          TextChannel
	ChannelID        string   `json:"channel_id"`
	Author           User     `json:"author"`
	Content          string   `json:"content"`
	Timestamp        string   `json:"timestamp"` // ISO8601 timestamp
	MentionsEveryone bool     `json:"mention_everyone"`
	UsersMentions    []User   `json:"mentions"`
	RolesMentions    []string `json:"mention_roles"`
	// TODO: Fix this
	// ChannelMentions  []ChannelMention `json:"mention_channels"`
	Reactions []Reaction `json:"reactions"`
	Embeds    []Embed    `json:"embeds"`
	Pinned    bool       `json:"pinned"`
	Type      int        `json:"type"`
	Flags     int        `json:"flags"`
}
type MessageData struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
	Flags   int     `json:"flags"`
}
type payloadMessage struct {
	Content   string                   `json:"content"`
	Embeds    []Embed                  `json:"embeds"`
	Reference *payloadMessageReference `json:"message_reference,omitempty"`
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
			Reference: &payloadMessageReference{
				ID:   m.ID,
				Type: 0,
			},
		}

		var payload bytes.Buffer
		json.NewEncoder(&payload).Encode(message)

		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages", API_URL, m.ChannelID), &payload)
		request.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		request.Header.Add("Content-Type", "application/json")
		http.DefaultClient.Do(request)
	case MessageData:
		message := payloadMessage{
			Content: data.Content,
			Embeds:  data.Embeds,
			Reference: &payloadMessageReference{
				ID:   m.ID,
				Type: 0,
			},
		}

		var payload bytes.Buffer
		json.NewEncoder(&payload).Encode(message)
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages", API_URL, m.ChannelID), &payload)
		request.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		request.Header.Add("Content-Type", "application/json")
		http.DefaultClient.Do(request)
	}
}

func (m Message) Delete(Client Client) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s", API_URL, m.ChannelID, m.ID), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 204 {
		panic("Error deleting message")
	}
}

func (m Message) React(Client Client, Emoticon any) {
	switch emoji := Emoticon.(type) {
	case rune:
		if !IsEmoji(emoji) {
			panic("Invalid emoji")
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", API_URL, m.ChannelID, m.ID, url.PathEscape(string(emoji))), nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 204 {
			panic(fmt.Sprintf("An error has occured. Status code: %d", res.StatusCode))
		}

	case string:
		if !IsCustomEmoji(emoji) {
			panic("Invalid emoji")
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", API_URL, m.ChannelID, m.ID, url.PathEscape(strings.Split(emoji, ":")[1]+":"+strings.TrimSuffix(strings.Split(emoji, ":")[2], ">"))), nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 204 {
			panic(fmt.Sprintf("An error has occured. Status code: %d", res.StatusCode))
		}
	case Emoji:
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", API_URL, m.ChannelID, m.ID, url.PathEscape(fmt.Sprintf("%s:%s", emoji.Name, emoji.ID))), nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 204 {
			panic(fmt.Sprintf("An error has occured. Status code: %d", res.StatusCode))
		}
	default:
		panic("Cannot react with another type than: rune (real emoji), string (custom emoji) or emoji object.")
	}
}
func (m Message) RemoveReact(Client Client, Emoticon any) {
	switch emoji := Emoticon.(type) {
	case rune:
		if !IsEmoji(emoji) {
			panic("Invalid emoji")
		}
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", API_URL, m.ChannelID, m.ID, url.PathEscape(string(emoji))), nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 204 {
			panic(fmt.Sprintf("An error has occured. Status code: %d", res.StatusCode))
		}

	case string:
		if !IsCustomEmoji(emoji) {
			panic("Invalid emoji")
		}
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", API_URL, m.ChannelID, m.ID, url.PathEscape(strings.Split(emoji, ":")[1]+":"+strings.TrimSuffix(strings.Split(emoji, ":")[2], ">"))), nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 204 {
			panic(fmt.Sprintf("An error has occured. Status code: %d", res.StatusCode))
		}
	case Emoji:
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", API_URL, m.ChannelID, m.ID, url.PathEscape(fmt.Sprintf("%s:%s", emoji.Name, emoji.ID))), nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bot %s", Client.Token))
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 204 {
			panic(fmt.Sprintf("An error has occured. Status code: %d", res.StatusCode))
		}
	default:
		panic("Cannot remove reaction with another type than: rune (real emoji), string (custom emoji) or emoji object.")
	}
}

func IsEmoji(char rune) bool {
	for _, rt := range emojiRanges {
		if unicode.Is(&rt, char) {
			return true
		}
	}
	return false
}

func IsCustomEmoji(str string) bool {
	reg := regexp.MustCompile(`^<a?:\w+:\d+>$`)
	return reg.MatchString(str)
}
