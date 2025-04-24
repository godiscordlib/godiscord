package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

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
		message_args := strings.Fields(message.Content)
		if message.Author.Bot {
			return
		}
		if strings.HasPrefix(message.Content, "!bd") {
			if len(message_args) < 2 || !IsNumber(message_args[1]) {
				message.Reply(Client, "Wrong usage. Please provide a real integer.")
				return
			}
			bd_len, err := strconv.Atoi(message_args[1])
			if err != nil {
				message.Reply(Client, "An error has occured while converting str to int.")
				return
			}
			message.Channel.BulkDelete(Client, bd_len)
			msg := message.Channel.Send(Client, fmt.Sprintf("Deleted %d messages", bd_len))
			msg.Edit(Client, "Ezzzzzz")
		} else {
			// message.Reply(Client, common.MessageData{
			// 	Components: []common.ActionRow{
			// 		common.NewActionRow().AddComponent(common.Button{
			// 			Type:     2,
			// 			Style:    enums.Success,
			// 			CustomID: "hello",
			// 			Label:    "Hi",
			// 			Emoji:    "😀",
			// 		}),
			// 	},
			// })
		}

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

func IsNumber(str string) bool {
	for _, char := range str {
		if !unicode.IsNumber(char) {
			return false
		}
	}
	return true
}
