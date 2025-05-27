package main

import (
	"fmt"
	"os"
	"strings"

	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
)

func main() {
	token, err := os.ReadFile("token.txt")
	if err != nil {
		panic(err)
	}
	Client := classes.NewClient(
		enums.GatewayIntent.Guilds,
		enums.GatewayIntent.GuildMessages,
		enums.GatewayIntent.MessageContent,
	)

	Client.On("READY", func(args ...any) {
		c := args[0].(classes.Client)
		fmt.Println(c.Username, "is ready")
		err := c.SetPresence(classes.PresenceUpdate{
			Activities: []classes.Activity{
				{
					Name: "godiscord",
					Type: enums.ActivityType.Watching,
				},
			},
			Status: "idle",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		c.GetWSPing()
	})
	Client.On("MESSAGE_CREATE", func(args ...any) {
		// message := args[0].(classes.Message)
		// if message.Content == "!sf" {
		// 	now := time.Now()
		// 	message.Reply(classes.MessageData{
		// 		Content: "Hello",
		// 		Attachments: []classes.Attachment{
		// 			{
		// 				FileName:    "godiscord.webp",
		// 				FilePath:    "./www/public/godiscord.webp",
		// 				Description: "Hello",
		// 			},
		// 		},
		// 		Embeds: []classes.Embed{
		// 			classes.NewEmbed().SetThumbnail("attachment://godiscord.webp", 64, 64),
		// 		},
		// 	})
		// 	timeItTook := time.Since(now).Milliseconds()
		// 	message.Reply(strconv.Itoa(int(timeItTook)))
		// }
		message := args[0].(classes.Message)
		//var message_args []string
		commandName := strings.Fields(message.Content)[0]
		if len(strings.Fields(message.Content)) > 1 {
			//message_args = strings.Fields(message.Content)[1:]
		}
		if commandName == "!ping" {
			message.Reply(fmt.Sprintf("Pong!\n%dms", Client.WS.Ping))
		}
		if commandName == "!gh" || commandName == "!github" {
			message.Reply("https://github.com/AYn0nyme/godiscord")
		}
	})
	err = Client.Connect(strings.TrimSpace(string(token)))
	if err != nil {
		fmt.Println(err)
		return
	}
}
