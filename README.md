# ⚠️ Project still under construction, bugs may occur.

<div align="center">
    <img src="./www/static/logo@128x128.webp"/>
    <h1>godiscord</h1>
</div>

`godiscord` is a lightweight library for interacting with the Discord API in Go, inspired by [discord.js](https://discord.js.org). It provides a simple and intuitive interface to create powerful and performant Discord bots in Go.

## 📦 Installation

```bash
go get github.com/AYn0nyme/godiscord
```

## 🧑‍💻 Example usage

Here is a complete example of a Discord bot in Go using `godiscord`:

```go
package main

import (
    "fmt"
    "os"
    "github.com/AYn0nyme/godiscord/client"
    "github.com/AYn0nyme/godiscord/internal/common"
    "github.com/AYn0nyme/godiscord/internal/enums"
)

func main() {
    Token, err := os.ReadFile("token.txt")
    if err != nil {
        panic(err)
    }
    Client := client.NewClient(string(Token), enums.GI_MessageContent, enums.GI_Guilds, enums.GI_GuildMessages)
    
    // READY event
    Client.On("READY", func(args ...any) {
        c := args[0].(common.Client)
        fmt.Println(c.Username, "is ready")
    })
    
    // React to messages
    Client.On("MESSAGE_CREATE", func(args ...any) {
        Message := args[0].(common.Message)
        if Message.Author.Bot {
            return
        }
        Message.React('🧙')
    })
    
    // Connect to Discord
    Client.Connect()
}
```

## 🛠️ Features

- **Connect to Discord** using a bot token.
- **Message handling**: send and receive messages.
- **Webhooks**, **embeds**, and more to come.

## 📜 License

This project is licensed under the MIT License. See the `LICENSE` file for details.
