package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	*User
	Base
	*EventManager
	Token   string
	Intents int
}

func (c Client) Connect() {
	websocket := WebSocket{}
	websocketChannel := make(chan WebSocketPayload)
	go websocket.Connect(c.Token, c.Intents, websocketChannel)
	for payload := range websocketChannel {
		switch payload.EventName {
		case "READY":
			var thing map[string]any
			err := json.Unmarshal(payload.Data, &thing)
			if err != nil {
				fmt.Println(err)
				return
			}
			var userData map[string]any
			userData = thing["user"].(map[string]any)
			c.User = &User{
				Base: Base{
					ID: toString(userData["id"]),
				},
				Username:      toString(userData["username"]),
				Discriminator: toString(userData["discriminator"]),
				AvatarHash:    toStringPtr(userData["avatar"]),
				Bot:           true,
				Global_Name:   toStringPtr(userData["global_name"]),
				Flags:         toIntPtr(userData["flags"]),
				VerifiedBot:   toBoolPtr(userData["verified"]),
			}
			c.Emit("READY", c)
		case "MESSAGE_CREATE":
			var message Message
			json.Unmarshal(payload.Data, &message)
			ptr_channel, err := c.GetChannelByID(message.ChannelID)
			if err != nil {
				panic(err)
			}
			if ptr_channel == nil {
				ptr_channel = &TextChannel{}
			}
			message.Channel = *ptr_channel
			c.Emit("MESSAGE_CREATE", message)
		case "MESSAGE_UPDATE":
			var message Message
			json.Unmarshal(payload.Data, &message)
			ptr_channel, err := c.GetChannelByID(message.ChannelID)
			if err != nil {
				panic(err)
			}
			if ptr_channel == nil {
				ptr_channel = &TextChannel{}
			}
			message.Channel = *ptr_channel
			c.Emit("MESSAGE_EDIT", message)
		case "MESSAGE_REACTION_ADD":
			var message Message
			json.Unmarshal(payload.Data, &message)
			ptr_channel, err := c.GetChannelByID(message.ChannelID)
			if err != nil {
				panic(err)
			}
			if ptr_channel == nil {
				ptr_channel = &TextChannel{}
			}
			message.Channel = *ptr_channel
			c.Emit("MESSAGE_REACTION_ADD", message)
		default:
			fmt.Println(payload.EventName)
		}
	}
}

func (c Client) GetChannelByID(ID string) (*TextChannel, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/channels/%s", API_URL, ID), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bot %s", c.Token))
	res, err := http.DefaultClient.Do(request)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	body_in_bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var channel TextChannel
	json.Unmarshal(body_in_bytes, &channel)
	return &channel, nil
}

func toString(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	}
	if num, ok := value.(float64); ok {
		return fmt.Sprintf("%.0f", num)
	}
	return ""
}

func toStringPtr(value interface{}) *string {
	if str, ok := value.(string); ok {
		return &str
	}
	return nil
}

func toIntPtr(value interface{}) *int {
	if num, ok := value.(float64); ok {
		n := int(num)
		return &n
	}
	return nil
}

func toBoolPtr(value interface{}) *bool {
	if b, ok := value.(bool); ok {
		return &b
	}
	return nil
}
