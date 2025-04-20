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
	Client := client.NewClient(string(Token), 33281)
	Client.On("READY", func(args ...interface{}) {
		c := args[0].(common.Client)
		fmt.Println(c.Username, "is ready")
	})
	Client.On("MESSAGE_CREATE", func(args ...interface{}) {
		Message := args[0].(common.Message)
		if !Message.Author.Bot {
			Message.Reply(Client, common.MessageData{
				Embeds: []common.Embed{
					common.NewEmbed().SetTitle("Hey").SetDescription("C'était rapide à implémenter").SetFooter("et pourquoi pas un footer", "https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png").SetURL("https://archlinux.org").AddField("un", "field aussi", false),
				},
			})
		}
	})
	Client.Connect()
}
