package main

import (
	"fmt"
	"os"
	"strings"

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
		message := args[0].(common.Message)
		var message_args []string
		if len(strings.Fields(message.Content)) > 1 {
			message_args = strings.Fields(message.Content)[1:]
		}
		if message.Author.Bot {
			return
		}
		if strings.HasPrefix(message.Content, "!ban") {
			if len(message_args) < 1 {
				message.Reply(Client, "noob not enof arguments")
			}
			if len(message.UsersMentions) > 0 {
				err := message.Channel.Guild.Ban(Client, message.UsersMentions[0], 0)
				if err != nil {
					panic(err)
				}
			}
		}
	})
	Client.Connect()
}
