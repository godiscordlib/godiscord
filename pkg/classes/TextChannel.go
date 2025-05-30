package classes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"godiscord.foo.ng/lib/pkg/enums"
)

type TextChannel struct {
	Channel
}

// Reply sends a message in the text channel
func (t Channel) Reply(data any) (*Message, error) {
	var req *http.Request
	var contentType string
	var body io.Reader
	var message Message

	switch v := data.(type) {
	case string:
		payload := payloadMessage{
			Content: v,
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

			var flags int
			for _, flag := range v.Flags {
				flags += int(flag)
			}

			msg := payloadMessage{
				Content:     v.Content,
				Embeds:      v.Embeds,
				Components:  v.Components,
				Flags:       flags,
				Attachments: v.Attachments,
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

	url := fmt.Sprintf("%s/channels/%s/messages", API_URL, t.ID)
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

func (t Channel) BulkDelete(Messages any) error {
	if t.Type != enums.ChannelType.TextChannel {
		return errors.New("error: wrong channel type")
	}
	switch messages_for_req := Messages.(type) {
	case int:
		get_messages_req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/channels/%s/messages?limit=%d", API_URL, t.ID, messages_for_req), nil)
		if err != nil {
			return err
		}
		get_messages_req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
		var res_messages []Message
		get_messages_res, err := http.DefaultClient.Do(get_messages_req)
		if err != nil {
			return err
		}
		defer get_messages_res.Body.Close()
		get_messages_body, err := io.ReadAll(get_messages_res.Body)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(get_messages_body, &res_messages); err != nil {
			return err
		}
		var messages_ids []string
		for _, v := range res_messages {
			messages_ids = append(messages_ids, v.ID)
		}
		messages_for_req_ := map[string]any{
			"messages": messages_ids,
		}
		req_body_bytes, err := json.Marshal(messages_for_req_)
		if err != nil {
			return err
		}
		req_body := bytes.NewReader(req_body_bytes)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages/bulk-delete", API_URL, t.ID), req_body)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != 204 {
			return fmt.Errorf("error: got status code %d while bulk deleting instead of 204", res.StatusCode)
		}
	case []string:
		req_data := fmt.Sprintf(`{"messages":[%s]}`, strings.Join(messages_for_req, ","))
		req_body := bytes.NewReader([]byte(req_data))
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages/bulk-delete", API_URL, t.ID), req_body)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != 204 {
			return fmt.Errorf("error: got status code %d while bulk deleting instead of 204", res.StatusCode)
		}
	case []Message:
		var messages_ids []string
		for _, v := range messages_for_req {
			messages_ids = append(messages_ids, fmt.Sprintf(`"%s"`, v.ID))
		}
		req_data := fmt.Sprintf(`{"messages":[%s]}`, strings.Join(messages_ids, ","))
		req_body := bytes.NewReader([]byte(req_data))
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages/bulk-delete", API_URL, t.ID), req_body)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 204 {
			return fmt.Errorf("error: got status code %d while bulk deleting instead of 204", res.StatusCode)
		}
	default:
		return fmt.Errorf("error: wrong type using bulk delete")
	}
	return nil
}
