package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
	"github.com/godiscordlib/godiscord/pkg/new"
	"github.com/godiscordlib/godiscord/pkg/slash"
	"github.com/godiscordlib/godiscord/pkg/types"
	"github.com/godiscordlib/godiscord/pkg/utils"
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

	// err = slash.RegisterGuildCommands("1375914465064915144", []classes.SlashCommandData{
	// 	{
	// 		Name:                      "say",
	// 		DefaultMembersPermissions: []types.Permission{enums.Permission.SendMessages},
	// 		Description:               "repeat what you say",
	// 		Options: []classes.SlashCommandOptionInt{
	// 			classes.SlashCommandStringOption{
	// 				SlashCommandOption: classes.SlashCommandOption{
	// 					Name:        "msg",
	// 					Description: "message to send",
	// 					Type:        enums.SlashCommandOptionType.String,
	// 					Required:    true,
	// 				},
	// 			},
	// 		},
	// 	},
	// }, "1373794354677813290")

	err = slash.RegisterGuildCommands(
		"1375914465064915144",
		[]classes.SlashCommandData{
			new.SlashCommand().
				SetName("say").
				SetDefaultMembersPermissions(enums.Permission.SendMessages).
				SetDescription("repeat what you say").
				AddOption(
					new.SlashCommandStringOption().
						SetName("msg").
						SetDescription("message to send").
						SetRequired(true),
				),
		},
		"1373794354677813290",
	)
	Client.On("GUILD_CREATE", func(args ...any) {
		fmt.Println(args[0].(classes.Guild).Name)
	})

	Client.On("READY", func(args ...any) {
		c := args[0].(*classes.Client)
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
			message.Reply(fmt.Sprintf("Pong!\n%dms", Client.GetWSPing()))
		}
		if commandName == "!gh" || commandName == "!github" {
			err = message.Reply(classes.MessageData{
				Components: []classes.ActionRow{
					new.ActionRow().AddComponent(new.RoleSelectMenu().SetCustomID("hello")),
				},
			})
			fmt.Println(err)
		}
	})
	Client.On("INTERACTION_CREATE", func(args ...any) {
		interaction := args[0].(classes.BaseInteraction)
		if interaction.Type == enums.InteractionResponseType.ApplicationCommand {
			if interaction.GetName() == "ping" {
				interaction.Reply(classes.MessageData{
					Embeds: []classes.Embed{
						new.Embed().SetDescription(fmt.Sprintf("🏓 %dms", Client.GetWSPing())).SetColor("00ADD8"),
					},
					Components: []classes.ActionRow{
						new.ActionRow().AddComponent(new.Button().SetCustomID("rori").SetLabel("Clik").SetStyle(enums.ButtonType.Success)),
					},
				})
			} else if interaction.GetName() == "say" {
				tmp := *interaction.Data.Options
				if tmp == nil {
					return
				}
				msgPtr, _ := interaction.Reply(classes.MessageData{
					Content: "Sending message.",
					Flags:   []types.MessageFlag{enums.MessageFlags.Ephemeral},
				})
				interaction.Channel.(classes.TextChannel).Send(*tmp[0].Value)
				if msgPtr == nil {
					fmt.Println("u nob")
				}
				msg := *msgPtr
				msg.Edit("Sent message :white_check_mark:")
			} else {
				interaction.Reply(classes.MessageData{
					Content: ":x: Unknown command.",
					Flags:   []types.MessageFlag{enums.MessageFlags.Ephemeral},
				})
			}
		}
		if interaction.Type == enums.InteractionResponseType.MessageComponent {
			interaction.Reply(classes.MessageData{
				Embeds: []classes.Embed{
					new.Embed().SetDescription(utils.RoleString(interaction.Values()[0])).SetColor("00ADD8"),
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
