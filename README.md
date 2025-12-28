# ⚠️ Project still under construction, bugs may occur.

<p align="center" style="margin-bottom: 0px !important;">
  <img src="./www/public/godiscord.webp" width="144" height="144">
</p>
<h1 align="center" style="margin-top: 0px;">godiscord</h1>

`godiscord` is a lightweight library for interacting with the Discord API in Go, inspired by [discord.js](https://discord.js.org). It provides a simple and intuitive interface to create powerful and performant Discord bots in Go.

## 📦 Installation

```bash
go get github.com/godiscordlib/godiscord/
```

## 🧑‍💻 Example usage

Here is a complete example of a Discord bot in Go using `godiscord`:

```go
package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
	"github.com/godiscordlib/godiscord/pkg/slash"
	"github.com/godiscordlib/godiscord/pkg/new"
)

func main() {
	Client := new.Client(
		os.Getenv("DISCORD_TOKEN"),
		enums.GatewayIntent.Guilds,
		enums.GatewayIntent.GuildMembers,
		enums.GatewayIntent.GuildMessages,
		enums.GatewayIntent.GuildModeration,
		enums.GatewayIntent.MessageContent,
	)

	slash.RegisterGuildCommands("1375914465064915144", []classes.SlashCommandData{
		{
			Name:        "ping",
			Description: "Pong! Get the ping of the bot",
			Type:        enums.InteractionType.ChatInput,
		},
	}, "1373794354677813290")

	Client.On("READY", func(args ...any) {
		fmt.Println("READY:", args[0].(*classes.Client).User.Username)
		Client.SetPresence(classes.PresenceUpdate{
			Activities: []classes.Activity{
				classes.Activity{
					Name: "godiscord",
					Type: enums.ActivityType.Streaming,
					URL:  "https://twitch.tv/godiscord",
				},
			},
			Status: "dnd",
			AFK:    false,
		})
	})

	Client.On("INTERACTION_CREATE", func(args ...any) {
		interaction := args[0].(classes.BaseInteraction)
		if interaction.GetName() == "ping" {
			interaction.Reply(classes.MessageData{
				Embeds: []classes.Embed{
					new.Embed().SetDescription(fmt.Sprintf("🏓 **%d**ms", Client.GetWSPing())).SetColor("00ADD8"),
				},
			})
		}
	})

	Client.Connect()
}


```

## 🛠️ Features

- **Connect to Discord** using a bot token.
- **Message handling**: send and receive messages.
- **Webhooks**, **embeds**, and more to come.

## 🧑‍💻 Contribute

- First `git clone https://github.com/AYn0nyme/godiscord`
- Then `go mod tidy`
- And you're good to **Go** 😄

## 📜 License

This project is licensed under the MIT License. See the `LICENSE` file for details.
