package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AYn0nyme/godiscord/client"
	"github.com/AYn0nyme/godiscord/internal/common"
	"github.com/AYn0nyme/godiscord/internal/enums"
)

func main() {
	token, err := os.ReadFile("token.txt")
	if err != nil {
		panic(err)
	}
	Client := client.NewClient(string(token), enums.GI_MessageContent, enums.GI_Guilds, enums.GI_GuildMessages, enums.GI_GuildModeration)

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
				message.Reply("not enough arguments")
			}
			if len(message.UsersMentions) > 0 {
				err := message.Channel.Guild.Ban(message.UsersMentions[0], common.BanOptions{
					DeleteMessageSeconds: 0,
					Reason:               "because",
				})
				if err != nil {
					panic(err)
				}
			} else {
				err := message.Channel.Guild.Ban(message_args[0], common.BanOptions{
					DeleteMessageSeconds: 0,
					Reason:               "i can",
				})
				if err != nil {
					panic(err)
				}
			}
			message.Reply("Banned the user :)")
		} else if strings.HasPrefix(message.Content, "!cg") {
			guild, err := Client.CreateGuild(common.CreateGuildOptions{
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
			ch, err := message.Channel.Guild.CreateChannel(common.CreateChannelOptions{
				Name:       message_args[0],
				Type:       enums.TextChannel,
				CategoryID: message_args[1],
				Position:   1,
			})
			if err != nil {
				panic(err)
			}
			message.Reply("created channel.")
			if ch.Type == enums.TextChannel {
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
		} else if strings.HasPrefix(message.Content, "!cr") {
			message.Channel.Guild.CreateRole(common.CreateRoleOptions{})
		} else if strings.HasPrefix(message.Content, "!dr") {
			if len(message_args) < 1 {
				message.Reply("not enough arguments")
				return
			}
			err = message.Channel.Guild.DeleteRole(message_args[0], message_args[1:]...)
			if err != nil {
				message.Reply("ERROR")
				fmt.Println(err)
			} else {
				message.Reply("DELETED ROLE")
			}
		} else if strings.HasPrefix(message.Content, "!eg") {
			_, error := message.Channel.Guild.Edit(common.EditGuildOptions{
				Name: "Edited by GODISCORD!",
			})
			if error != nil {
				panic(err)
			}
		}
	})
	Client.On("GUILD_CREATE", func(args ...any) {
		guild := args[0].(common.Guild)
		fmt.Println(guild.Name, guild.ID)
	})
	Client.Connect()
}
