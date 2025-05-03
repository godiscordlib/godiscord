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
			} else {
				err := message.Channel.Guild.Ban(Client, message_args[0], 0)
				if err != nil {
					panic(err)
				}
			}
			message.Reply(Client, "Banned the user :)")
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
				message.Reply(Client, "You're not my owner bozo")
				return
			}
			fmt.Println("im here")
			err := Client.LeaveGuild(message_args[0])
			if err == nil {
				message.Reply(Client, "I left the server successfully")
			} else {
				panic(err)
			}
		} else if strings.HasPrefix(message.Content, "!unban") {
			if len(message_args) < 1 {
				message.Reply(Client, "noob not enof arguments")
			}
			err := message.Channel.Guild.UnBan(Client, message_args[0], 0)
			if err != nil {
				panic(err)
			}
			message.Reply(Client, "Unbanned user :(")
		} else if strings.HasPrefix(message.Content, "!eg") {
			if message.Author.ID != "943580965446512661" {
				message.Reply(Client, "You're not my owner bozo")
				return
			}
			_, err := message.Channel.Guild.Edit(Client, common.EditGuildOptions{
				Name: message_args[0],
			})
			if err != nil {
				panic(err)
			}
			message.Reply(Client, "modified your server oboz")
		}
	})
	Client.On("GUILD_CREATE", func(args ...any) {
		guild := args[0].(common.Guild)
		fmt.Println(guild.Name, guild.ID)
	})
	Client.Connect()
}
