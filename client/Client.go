package client

import "github.com/AYn0nyme/godiscord/internal/common"

func NewClient(Token string, Intents int) common.Client {
	return common.Client{
		Token:        Token,
		Intents:      Intents,
		User:         &common.User{},
		EventManager: common.NewEventManager(),
	}
}
