package classes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"godiscord.foo.ng/lib/pkg/enums"
)

type interactionPayloadMessage struct {
	Type int                           `json:"type"`
	Data interactionPayloadMessageData `json:"data"`
}

type interactionPayloadMessageData struct {
	Content     string                   `json:"content"`
	Embeds      []Embed                  `json:"embeds,omitempty"`
	Flags       int                      `json:"flags,omitempty"`
	Components  []ActionRow              `json:"components,omitempty"`
	Attachments []Attachment             `json:"attachment,omitempty"`
	Files       []string                 `json:"files,omitempty"`
	Reference   *payloadMessageReference `json:"message_reference,omitempty"`
}

func (bi BaseInteraction) Reply(data any) (*Message, error) {
	if bi.Type != enums.InteractionResponseType.ApplicationCommand {
		return nil, nil
	}
	var req *http.Request
	var contentType string
	var body io.Reader
	var message Message

	switch v := data.(type) {
	case string:
		payload := interactionPayloadMessage{
			Type: 4,
			Data: interactionPayloadMessageData{
				Content: v,
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

			for i, path := range v.Files {
				file, err := os.Open(path)
				if err != nil {
					pw.CloseWithError(err)
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

			msg := interactionPayloadMessage{
				Type: 4,
				Data: interactionPayloadMessageData{
					Content:     v.Content,
					Embeds:      v.Embeds,
					Components:  v.Components,
					Flags:       v.Flags,
					Attachments: v.Attachments,
				},
			}
			jsonData, err := json.Marshal(msg)
			if err != nil {
				pw.CloseWithError(err)
				return
			}
			err = writer.WriteField("payload_json", string(jsonData))
			if err != nil {
				return
			}
		}()

		body = pr
		contentType = writer.FormDataContentType()
	default:
		return nil, fmt.Errorf("unsupported reply data type: %T", data)
	}

	url := fmt.Sprintf("%s/interactions/%s/%s/callback", API_URL, bi.ID, bi.Token)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("Content-Type", contentType)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resp, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to reply: %d - %s", res.StatusCode, string(resp))
	}
	if err = json.Unmarshal(resp, &message); err != nil {
		return nil, err
	}
	return &message, nil
}
