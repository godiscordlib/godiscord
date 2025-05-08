package client

import (
	"github.com/AYn0nyme/godiscord/internal/common"
	"github.com/AYn0nyme/godiscord/internal/enums"
)

func NewClient(Token string, Intents ...enums.GatewayIntent) common.Client {
	var intents int
	for _, intent := range Intents {
		intents += int(intent)
	}
	return common.Client{
		Token:        Token,
		Intents:      intents,
		User:         &common.User{},
		EventManager: common.NewEventManager(),
	}
}
