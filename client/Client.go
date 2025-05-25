package client

import (
	"os"

	"godiscord.foo.ng/lib/internal/common"
	"godiscord.foo.ng/lib/internal/enums"
)

func NewClient(Token string, Intents ...enums.GatewayIntent) common.Client {
	var intents int
	for _, intent := range Intents {
		intents += int(intent)
	}
	os.Setenv("GODISCORD_TOKEN", Token)
	return common.Client{
		Token:        Token,
		Intents:      intents,
		User:         &common.User{},
		EventManager: common.NewEventManager(),
		ReadyChan:    make(chan struct{}),
	}
}
