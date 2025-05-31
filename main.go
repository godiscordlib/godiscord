package main

import (
	"fmt"
	"os"
	"strings"

	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
	"godiscord.foo.ng/lib/pkg/new"
)

func main() {
	token, err := os.ReadFile("token.txt")
	if err != nil {
		panic(err)
	}
	Client := new.Client(
		strings.TrimSpace(string(token)),
		enums.GatewayIntent.Guilds,
		enums.GatewayIntent.GuildMessages,
		enums.GatewayIntent.MessageContent,
	)

	Client.On("GUILD_CREATE", func(args ...any) {
		fmt.Println(args[0].(classes.Guild).Name)
	})

	Client.On("READY", func(args ...any) {
		c := args[0].(classes.Client)
		fmt.Println(c.Username+"#"+c.Discriminator, "is ready")
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
		if len(strings.Fields(message.Content)) <= 0 {
			return
		}
		//var message_args []string
		commandName := strings.Fields(message.Content)[0]
		if len(strings.Fields(message.Content)) > 1 {
			//message_args = strings.Fields(message.Content)[1:]
		}
		if commandName == "!ping" {
			message.Reply(fmt.Sprintf("Pong!\n%dms", Client.WS.Ping))
		}
		if commandName == "!gh" || commandName == "!github" {
			err = message.Reply(classes.MessageData{
				Components: []classes.ActionRow{
					new.ActionRow().AddComponent(classes.NewRoleSelectMenu().SetCustomID("hello")),
				},
			})
			fmt.Println(err)
		}
	})
	Client.On("INTERACTION_CREATE", func(args ...any) {
		interaction := args[0].(classes.BaseInteraction)
		if interaction.Type == enums.InteractionResponseType.ApplicationCommand {
			interaction.Reply(classes.MessageData{
				Embeds: []classes.Embed{
					new.Embed().SetDescription(fmt.Sprintf("🏓 %dms", Client.GetWSPing())).SetColor("00ADD8"),
				},
			})
		}
	})
	err = Client.Connect()

	if err != nil {
		fmt.Println(err)
		return
	}
}
