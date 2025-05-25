package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"godiscord.foo.ng/lib/internal/enums"
)

type TextChannel struct {
	BaseChannel
}

// Sends a message in the textchannel
func (t BaseChannel) Send(Data any) (*Message, error) {
	if t.Type != enums.TextChannel {
		return nil, errors.New("error: wrong channel type")
	}
	switch data := Data.(type) {
	case string:
		message := payloadMessage{
			Content: data,
		}

		var payload bytes.Buffer
		json.NewEncoder(&payload).Encode(message)

		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages", API_URL, t.ID), &payload)
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
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%s/messages", API_URL, t.ID), &payload)
		request.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
		request.Header.Set("Content-Type", "application/json")
		http.DefaultClient.Do(request)
	}
	get_messages_req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/channels/%s/messages?limit=%d", API_URL, t.ID, 1), nil)
	if err != nil {
		return nil, err
	}
	get_messages_req.Header.Set("Authorization", fmt.Sprintf("Bot %s", os.Getenv("GODISCORD_TOKEN")))
	var res_message []Message
	get_messages_res, err := http.DefaultClient.Do(get_messages_req)
	if err != nil {
		return nil, err
	}
	defer get_messages_res.Body.Close()
	get_messages_body, err := io.ReadAll(get_messages_res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(get_messages_body, &res_message); err != nil {
		return nil, err
	}
	return &res_message[0], nil
}

func (t BaseChannel) BulkDelete(Messages any) error {
	if t.Type != enums.TextChannel {
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
