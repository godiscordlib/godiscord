package main

import (
	"fmt"
	"os"

	"github.com/AYn0nyme/godiscord/client"
	"github.com/AYn0nyme/godiscord/internal/common"
)

func main() {
	Token, err := os.ReadFile("token.txt")
	if err != nil {
		panic(err)
	}
	Client := client.NewClient(string(Token), 34305)
	Client.On("READY", func(args ...any) {
		c := args[0].(common.Client)
		fmt.Println(c.Username, "is ready")
	})
	Client.On("MESSAGE_CREATE", func(args ...any) {
		Message := args[0].(common.Message)
		if Message.Author.Bot {
			return
		}
		Message.React(Client, '🧙')
	})
	Client.On("MESSAGE_EDIT", func(args ...any) {
		Message := args[0].(common.Message)
		if Message.Author.Bot {
			return
		}
		Message.RemoveReact(Client, '🧙')
	})
	Client.Connect()
}
