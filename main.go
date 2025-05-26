package main

import (
	"fmt"
	"os"
	"strings"

	"godiscord.foo.ng/lib/internal/types"
	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
)

func main() {
	token, err := os.ReadFile("token.txt")
	if err != nil {
		panic(err)
	}
	Client := classes.Client{
		Intents: []types.GatewayIntent{
			enums.GatewayIntent.Guilds,
			enums.GatewayIntent.GuildMessages,
			enums.GatewayIntent.MessageContent,
		},
	}

	Client.On("READY", func(args ...any) {
		c := args[0].(*classes.Client)
		fmt.Println(c.Username, "is ready")
	})
	Client.On("MESSAGE_CREATE", func(args ...any) {
		message := args[0].(classes.Message)
		message.Reply("Hi!")
	})
	Client.Connect(strings.TrimSpace(string(token)))
}
