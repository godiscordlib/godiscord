package client

import (
	"os"

	"godiscord.foo.ng/lib/internal/types"
	"godiscord.foo.ng/lib/pkg/classes"
)

func NewClient(Token string, Intents ...types.GatewayIntent) classes.Client {
	var intents int
	for _, intent := range Intents {
		intents += int(intent)
	}
	os.Setenv("GODISCORD_TOKEN", Token)
	return classes.Client{
		Token:        Token,
		Intents:      intents,
		User:         &classes.User{},
		EventManager: classes.NewEventManager(),
		ReadyChan:    make(chan struct{}),
	}
}
