package classes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	ChannelID        string      `json:"channel_id"`
	Author           User        `json:"author"`
	Content          string      `json:"content"`
	Timestamp        string      `json:"timestamp"` // ISO8601 timestamp
	MentionsEveryone bool        `json:"mention_everyone"`
	UsersMentions    []User      `json:"mentions"`
	RolesMentions    []string    `json:"mention_roles"`
	Components       []ActionRow `json:"components"`
	Reactions        []Reaction  `json:"reactions"`
	Embeds           []Embed     `json:"embeds"`
	Pinned           bool        `json:"pinned"`
	Type             int         `json:"type"`
	Flags            int         `json:"flags"`
	// TODO: Fix this
	// ChannelMentions  []ChannelMention `json:"mention_channels"`
}
type MessageData struct {
	Content     string       `json:"content"`
	Embeds      []Embed      `json:"embeds,omitempty"`
	Flags       int          `json:"flags,omitempty"`
	Components  []ActionRow  `json:"components,omitempty"`
	Attachments []Attachment `json:"attachment,omitempty"`
	Files       []string     `json:"files,omitempty"`
}
type payloadMessage struct {
	Content     string                   `json:"content"`
	Embeds      []Embed                  `json:"embeds,omitempty"`
	Flags       int                      `json:"flags,omitempty"`
	Components  []ActionRow              `json:"components,omitempty"`
	Attachments []Attachment             `json:"attachment,omitempty"`
	Files       []string                 `json:"files,omitempty"`
	Reference   *payloadMessageReference `json:"message_reference,omitempty"`
}
type payloadMessageReference struct {
	ID   string `json:"message_id"`
	Type int    `json:"type"`
}

func (m Message) Reply(data any) error {
	var req *http.Request
	var contentType string
	var body io.Reader

	switch v := data.(type) {
	case string:
		payload := payloadMessage{
			Content: v,
			Reference: &payloadMessageReference{
				ID:   m.ID,
				Type: 0,
			},
		}
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(payload)
		body = buf
		contentType = "application/json"

	case MessageData:
		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)

		go func() {
			defer pw.Close()
			defer writer.Close()

			// 📎 Ajout des fichiers simples
			for i, path := range v.Files {
				file, err := os.Open(path)
				if err != nil {
					pw.CloseWithError(err)
					return
				}
				part, err := writer.CreateFormFile(fmt.Sprintf("files[%d]", i), filepath.Base(path))
				if err != nil {
					file.Close()
					pw.CloseWithError(err)
					return
				}
				_, err = io.Copy(part, file)
				file.Close()
				if err != nil {
					pw.CloseWithError(err)
					return
				}
			}

			// 📎 Ajout des attachments (préremplis)
			for i := range v.Attachments {
				v.Attachments[i].ID = i
				file, err := os.Open(v.Attachments[i].FilePath)
				if err != nil {
					pw.CloseWithError(err)
					return
				}
				part, err := writer.CreateFormFile(fmt.Sprintf("files[%d]", i), v.Attachments[i].FileName)
				if err != nil {
					file.Close()
					pw.CloseWithError(err)
					return
				}
				_, err = io.Copy(part, file)
				file.Close()
				if err != nil {
					pw.CloseWithError(err)
					return
				}
			}

			// 📝 payload_json
			msg := payloadMessage{
				Content:     v.Content,
				Embeds:      v.Embeds,
				Components:  v.Components,
				Flags:       v.Flags,
				Attachments: v.Attachments,
				Reference: &payloadMessageReference{
					ID:   m.ID,
					Type: 0,
				},
			}
			jsonData, err := json.Marshal(msg)
			if err != nil {
				pw.CloseWithError(err)
				return
			}
			writer.WriteField("payload_json", string(jsonData))
		}()

		body = pr
		contentType = writer.FormDataContentType()
	default:
		return fmt.Errorf("unsupported reply data type: %T", data)
	}

	url := fmt.Sprintf("%s/channels/%s/messages", API_URL, m.ChannelID)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("Content-Type", contentType)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		resp, _ := io.ReadAll(res.Body)
		return fmt.Errorf("failed to reply: %d - %s", res.StatusCode, string(resp))
	}
	return nil
}

func (m Message) Post() error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages/%s/crosspost", API_URL, m.ChannelID, m.ID), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("error: got status code %d while crossposting message", res.StatusCode)
	}
	return nil
}

func (m Message) Edit(Data any) {
	switch data := Data.(type) {
	case string:
		message := payloadMessage{
			Content: data,
		}

		var payload bytes.Buffer
		json.NewEncoder(&payload).Encode(message)

		request, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/channels/%s/messages/%s", API_URL, m.ChannelID, m.ID), &payload)
		request.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
		request.Header.Set("Content-Type", "application/json")
		http.DefaultClient.Do(request)
	case MessageData:
		message := payloadMessage{
			Content: data.Content,
			Embeds:  data.Embeds,
		}

		var payload bytes.Buffer
		json.NewEncoder(&payload).Encode(message)
		request, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/channels/%s/messages/%s", API_URL, m.ChannelID, m.ID), &payload)
		request.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
		request.Header.Set("Content-Type", "application/json")
		http.DefaultClient.Do(request)
	}
}

func (m Message) Delete(Reason ...string) {
	var reason string
	if len(Reason) > 0 {
		reason = Reason[0]
	}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s", API_URL, m.ChannelID, m.ID), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
	req.Header.Set("X-Audit-Log-Reason", reason)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 204 {
		panic("Error deleting message")
	}
}

func (m Message) React(Emoticon any) {
	switch emoji := Emoticon.(type) {
	case rune:
		if !IsEmoji(emoji) {
			panic("Invalid emoji")
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", API_URL, m.ChannelID, m.ID, url.PathEscape(string(emoji))), nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
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
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
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
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
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

func (m Message) RemoveReact(Emoticon any) {
	switch emoji := Emoticon.(type) {
	case rune:
		if !IsEmoji(emoji) {
			panic("Invalid emoji")
		}
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", API_URL, m.ChannelID, m.ID, url.PathEscape(string(emoji))), nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
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
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
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
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
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

func (m Message) RemoveAllReact() {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions", API_URL, m.ChannelID, m.ID), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
}

func (m Message) RemoveEmojiReact(Emoticon any) {
	switch emoji := Emoticon.(type) {
	case rune:
		if !IsEmoji(emoji) {
			panic("Invalid emoji")
		}
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", API_URL, m.ChannelID, m.ID, url.PathEscape(string(emoji))), nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
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
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s", API_URL, m.ChannelID, m.ID, url.PathEscape(strings.Split(emoji, ":")[1]+":"+strings.TrimSuffix(strings.Split(emoji, ":")[2], ">"))), nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
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
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s", API_URL, m.ChannelID, m.ID, url.PathEscape(fmt.Sprintf("%s:%s", emoji.Name, emoji.ID))), nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
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

func (m Message) RemoveReactFromUser(UserID string, Emoticon any) error {
	switch emoji := Emoticon.(type) {
	case rune:
		if !IsEmoji(emoji) {
			return fmt.Errorf("error: invalid emoji")
		}
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/%s", API_URL, m.ChannelID, m.ID, url.PathEscape(string(emoji)), UserID), nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
		if err != nil {
			return err
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != 204 {
			return fmt.Errorf("an error has occured. Status code: %d", res.StatusCode)
		}

	case string:
		if !IsCustomEmoji(emoji) {
			return fmt.Errorf("error: invalid custom emoji")
		}
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/%s", API_URL, m.ChannelID, m.ID, url.PathEscape(strings.Split(emoji, ":")[1]+":"+strings.TrimSuffix(strings.Split(emoji, ":")[2], ">")), UserID), nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
		if err != nil {
			return err
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != 204 {
			return fmt.Errorf("an error has occured. Status code: %d", res.StatusCode)
		}
	case Emoji:
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/%s", API_URL, m.ChannelID, m.ID, url.PathEscape(fmt.Sprintf("%s:%s", emoji.Name, emoji.ID)), UserID), nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
		if err != nil {
			return err
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != 204 {
			return fmt.Errorf("an error has occured. Status code: %d", res.StatusCode)
		}
	default:
		return fmt.Errorf("error: cannot remove reaction with another type than: rune (real emoji), string (custom emoji) or emoji object")
	}
	return nil
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

func attachFile(writer *multipart.Writer, fieldName, path string, optionalName ...string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	filename := filepath.Base(path)
	if len(optionalName) > 0 {
		filename = optionalName[0]
	}

	part, err := writer.CreateFormFile(fieldName, filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	return err
}

func doRequest(req *http.Request) error {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error: got status %d while replying to message", res.StatusCode)
	}
	return nil
}
