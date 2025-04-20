package common

import (
	"encoding/json"
	"fmt"
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
			var thing map[string]interface{}
			err := json.Unmarshal(payload.Data, &thing)
			if err != nil {
				fmt.Println(err)
				return
			}
			var userData map[string]interface{}
			userData = thing["user"].(map[string]interface{})
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
			c.Emit("MESSAGE_CREATE", message)
		default:
			fmt.Println(payload.EventName)
		}
	}
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
