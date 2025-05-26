package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"godiscord.foo.ng/lib/client"
	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
)

func main() {
	token, err := os.ReadFile("token.txt")
	if err != nil {
		panic(err)
	}
	Client := client.NewClient(strings.TrimSpace(string(token)), enums.GatewayIntent.MessageContent, enums.GatewayIntent.Guilds, enums.GatewayIntent.GuildMessages, enums.GatewayIntent.GuildModeration, enums.GatewayIntent.GuildPresences)

	Client.On("READY", func(args ...any) {
		c := args[0].(*classes.Client)
		fmt.Println(c.Username, "is ready")
		Client.SetPresence(classes.PresenceUpdate{
			Since:  time.Now().Unix(),
			Status: "online",
			AFK:    false,
			Activities: []classes.Activity{
				{
					Name: "the support",
					Type: enums.ActivityType.Listening,
				},
			},
		})
	})
	Client.On("MESSAGE_CREATE", func(args ...any) {
		message := args[0].(classes.Message)
		var message_args []string
		if len(strings.Fields(message.Content)) > 1 {
			message_args = strings.Fields(message.Content)[1:]
		}
		if message.Author.Bot {
			return
		}
		if strings.HasPrefix(message.Content, "!ban") {
			if len(message_args) < 1 {
				message.Reply("not enough arguments")
			}
			if len(message.UsersMentions) > 0 {
				err := message.Channel.Guild.Ban(message.UsersMentions[0], classes.BanOptions{
					DeleteMessageSeconds: 0,
					Reason:               "because",
				})
				if err != nil {
					panic(err)
				}
			} else {
				err := message.Channel.Guild.Ban(message_args[0], classes.BanOptions{
					DeleteMessageSeconds: 0,
					Reason:               "i can",
				})
				if err != nil {
					panic(err)
				}
			}
			message.Reply("Banned the user :)")
		} else if strings.HasPrefix(message.Content, "!cg") {
			guild, err := Client.CreateGuild(classes.CreateGuildOptions{
				Name: "Godiscord",
			})
			fmt.Println(guild)
			if err != nil {
				panic(err)
			}
			if guild == nil {
				panic("error: guild is nil")
			}
			fmt.Printf("Created guild %s with success!\n", guild.Name)
		} else if strings.HasPrefix(message.Content, "!lg") {
			if message.Author.ID != "943580965446512661" {
				message.Reply("You're not my owner")
				return
			}
			err := Client.LeaveGuild(message_args[0])
			if err == nil {
				message.Reply("I left the server successfully")
			} else {
				panic(err)
			}
		} else if strings.HasPrefix(message.Content, "!unban") {
			if len(message_args) < 1 {
				message.Reply("not enough arguments")
			}
			err := message.Channel.Guild.UnBan(message_args[0], "")
			if err != nil {
				panic(err)
			}
			message.Reply("Unbanned user :(")
		} else if strings.HasPrefix(message.Content, "!cc") {
			if message.Author.ID != "943580965446512661" {
				message.Reply("You're not my owner")
				return
			}
			if len(message_args) < 2 {
				message.Reply("not enough args")
				return
			}
			ch, err := message.Channel.Guild.CreateChannel(classes.CreateChannelOptions{
				Name:       message_args[0],
				Type:       enums.ChannelType.TextChannel,
				CategoryID: message_args[1],
				Position:   1,
			})
			if err != nil {
				panic(err)
			}
			message.Reply("created channel.")
			if ch.Type == enums.ChannelType.TextChannel {
				ch.Send("<@" + message.Author.ID + ">")
			}
		} else if strings.HasPrefix(message.Content, "!gb") {
			bans, err := message.Channel.Guild.GetBans()
			if bans == nil || *bans == nil || err != nil {
				panic(err)
			} else {
				for _, v := range *bans {
					message.Reply(v.User.Username)
				}
			}
		} else if strings.HasPrefix(message.Content, "!eg") {
			_, error := message.Channel.Guild.Edit(classes.EditGuildOptions{
				Name: "Edited by GODISCORD!",
			})
			if error != nil {
				panic(err)
			}
		}
	})
	Client.On("GUILD_CREATE", func(args ...any) {
		guild := args[0].(classes.Guild)
		fmt.Println(guild.Name, guild.ID)
	})
	Client.Connect()

	select {}
}
