package client

import (
	"encoding/json"
	"fmt"
	"godiscord/types"
	"godiscord/websockets"
)

type Client struct {
	*types.User
	types.Base
	*types.EventManager
	Token   string
	Intents int
}

func (c Client) Connect() {
	websocket := websockets.WebSocket{}
	websocketChannel := make(chan types.WebSocketPayload)
	go websocket.Connect(c.Token, c.Intents, websocketChannel)
	for payload := range websocketChannel {
		switch payload.EventName {
		case "READY":
			var thing map[string]interface{}
			err := json.Unmarshal(payload.Data, &thing)
			if err != nil {
				fmt.Println("Erreur de désérialisation:", err)
				return
			}
			var userData map[string]interface{}
			userData = thing["user"].(map[string]interface{})
			c.User = &types.User{
				Base: types.Base{
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
		}
	}
}

func NewClient(Token string, Intents int) Client {
	return Client{
		Token:        Token,
		Intents:      Intents,
		User:         &types.User{},
		EventManager: types.NewEventManager(),
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
