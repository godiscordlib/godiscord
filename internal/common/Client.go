package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AYn0nyme/godiscord/internal/enums"
)

type Client struct {
	*User
	Base
	*EventManager
	Token     string
	Intents   int
	ws        WebSocket
	wschannel chan webSocketPayload
	// guildCache map[string]Guild
}

// type guildMembersChunkEvent struct {
// 	GuildID string        `json:"guild_id"`
// 	Members []GuildMember `json:"members"`
// }

func (c Client) Connect() error {
	c.ws = WebSocket{}
	c.wschannel = make(chan webSocketPayload)
	go c.ws.Connect(c.Token, c.Intents, c.wschannel)
	for payload := range c.wschannel {
		switch payload.EventName {
		case "READY":
			var thing map[string]any
			err := json.Unmarshal(payload.Data, &thing)
			if err != nil {
				return err
			}
			var userData map[string]any
			userData = thing["user"].(map[string]any)
			c.User = &User{
				Base: Base{
					ID: toString(userData["id"]),
				},
				Username:      toString(userData["username"]),
				Discriminator: toString(userData["discriminator"]),
				AvatarHash:    toStringPtr(userData["avatar"]),
				Bot:           true,
				Global_Name:   toStringPtr(userData["global_name"]),
				Flags:         toIntPtr(userData["flags"]),
				VerifiedBot:   toBoolPtr(userData["verified"]),
			}
			c.Emit("READY", c)
		case "MESSAGE_CREATE":
			var message Message
			json.Unmarshal(payload.Data, &message)
			ptr_channel, err := c.GetTextChannelByID(message.ChannelID)
			if err != nil {
				return err
			}
			if ptr_channel == nil {
				ptr_channel = &TextChannel{}
			}
			message.Channel = *ptr_channel
			ptr_owner, err := message.Channel.Guild.GetMemberByID(message.Channel.Guild.OwnerID)
			if err != nil {
				return err
			}
			if ptr_owner == nil {
				ptr_owner = &GuildMember{}
			}
			message.Channel.Guild.Owner = *ptr_owner
			c.Emit("MESSAGE_CREATE", message)
		case "MESSAGE_UPDATE":
			var message Message
			fmt.Println(string(payload.Data))
			json.Unmarshal(payload.Data, &message)
			ptr_channel, err := c.GetTextChannelByID(message.ChannelID)
			if err != nil {
				return err
			}
			if ptr_channel == nil {
				ptr_channel = &TextChannel{}
			}
			message.Channel = *ptr_channel
			ptr_owner, err := message.Channel.Guild.GetMemberByID(message.Channel.Guild.OwnerID)
			if err != nil {
				return err
			}
			if ptr_owner == nil {
				ptr_owner = &GuildMember{}
			}
			message.Channel.Guild.Owner = *ptr_owner
			c.Emit("MESSAGE_UPDATE", message)
		case "MESSAGE_REACTION_ADD":
			var message Message
			json.Unmarshal(payload.Data, &message)
			ptr_channel, err := c.GetTextChannelByID(message.ChannelID)
			if err != nil {
				return err
			}
			if ptr_channel == nil {
				ptr_channel = &TextChannel{}
			}
			message.Channel = *ptr_channel
			ptr_owner, err := message.Channel.Guild.GetMemberByID(message.Channel.Guild.OwnerID)
			if err != nil {
				return err
			}
			if ptr_owner == nil {
				ptr_owner = &GuildMember{}
			}
			message.Channel.Guild.Owner = *ptr_owner
			c.Emit("MESSAGE_REACTION_ADD", message)
		case "GUILD_CREATE":
			var guild Guild
			json.Unmarshal(payload.Data, &guild)
			ptr_owner, err := guild.GetMemberByID(guild.OwnerID)
			if err != nil {
				return err
			}
			if ptr_owner == nil {
				ptr_owner = &GuildMember{}
			}
			guild.Owner = *ptr_owner
			// c.guildCache[guild.ID] = guild
			c.Emit("GUILD_CREATE", guild)
		case "GUILD_DELETE":
			var guild Guild
			json.Unmarshal(payload.Data, &guild)
			ptr_owner, err := guild.GetMemberByID(guild.OwnerID)
			if err != nil {
				return err
			}
			if ptr_owner == nil {
				ptr_owner = &GuildMember{}
			}
			guild.Owner = *ptr_owner
			c.Emit("GUILD_DELETE", guild)
		case "GUILD_UPDATE":
			var guild Guild
			json.Unmarshal(payload.Data, &guild)
			ptr_owner, err := guild.GetMemberByID(guild.OwnerID)
			if err != nil {
				return err
			}
			if ptr_owner == nil {
				ptr_owner = &GuildMember{}
			}
			guild.Owner = *ptr_owner
			c.Emit("GUILD_UPDATE", guild)
		case "GUILD_ROLE_CREATE":
			var role Role
			err := json.Unmarshal(payload.Data, &role)
			if err != nil {
				return err
			}
			c.Emit("GUILD_ROLE_CREATE", role)
		case "GUILD_ROLE_DELETE":
			var role Role
			err := json.Unmarshal(payload.Data, &role)
			if err != nil {
				return err
			}
			c.Emit("GUILD_ROLE_DELETE", role)
		case "GUILD_ROLE_UPDATE":
			var role Role
			err := json.Unmarshal(payload.Data, &role)
			if err != nil {
				return err
			}
			c.Emit("GUILD_ROLE_UPDATE", role)
		case "CHANNEL_CREATE":
			var channel BaseChannel
			err := json.Unmarshal(payload.Data, &channel)
			if err != nil {
				return err
			}
			c.Emit("CHANNEL_CREATE", channel)
		default:
			fmt.Println(payload.EventName)
		}
	}
	return nil
}

func (c Client) GetTextChannelByID(ID string) (*TextChannel, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/channels/%s", API_URL, ID), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bot "+c.Token)
	res, err := http.DefaultClient.Do(request)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	body_in_bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var channel TextChannel
	json.Unmarshal(body_in_bytes, &channel)
	if channel.Type != enums.TextChannel {
		return nil, fmt.Errorf("error: Channel is not a TextChannel")
	}

	guild, err := c.GetGuildByID(channel.GuildID)
	if err != nil {
		return nil, err
	}
	channel.Guild = *guild

	return &channel, nil
}

func (c Client) GetGuildByID(ID string) (*Guild, error) {
	var guild Guild
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/guilds/%s", API_URL, ID), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bot "+c.Token)
	res, err := http.DefaultClient.Do(request)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	body_in_bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	json.Unmarshal(body_in_bytes, &guild)
	return &guild, nil
}

// Can only be used if the bot is in less than 10 guilds.
func (c Client) CreateGuild(Options CreateGuildOptions) (*Guild, error) {
	body, err := json.Marshal(Options)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	body_reader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/guilds", API_URL), body_reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		bodyErr, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("status %d: %s", res.StatusCode, string(bodyErr))
	}
	defer res.Body.Close()
	body_in_bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var guild Guild
	if err = json.Unmarshal(body_in_bytes, &guild); err != nil {
		return nil, err
	}
	return &guild, err
}

func (c Client) LeaveGuild(GuildID string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/users/@me/guilds/%s", API_URL, GuildID), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bot "+c.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		fmt.Printf("%s\n", string(body))
		return fmt.Errorf("error: got status code %d while leaving guild", res.StatusCode)
	}
	return nil
}

// func (c Client) Edit()

// func (c Client) GetGuildMembers(Guild Guild) (*[]GuildMember, error) {
// 	var gms []GuildMember

// 	return &gms, nil
// }

func toString(value any) string {
	if str, ok := value.(string); ok {
		return str
	}
	if num, ok := value.(float64); ok {
		return fmt.Sprintf("%.0f", num)
	}
	return ""
}

func toStringPtr(value any) *string {
	if str, ok := value.(string); ok {
		return &str
	}
	return nil
}

func toIntPtr(value any) *int {
	if num, ok := value.(float64); ok {
		n := int(num)
		return &n
	}
	return nil
}

func toBoolPtr(value any) *bool {
	if b, ok := value.(bool); ok {
		return &b
	}
	return nil
}
