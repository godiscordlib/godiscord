package classes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/godiscordlib/godiscord/pkg/enums"
	"github.com/godiscordlib/godiscord/pkg/types"
)

const API_VERSION = "10"
const API_URL = "https://discord.com/api/v" + API_VERSION

type Client struct {
	*User
	*EventManager
	Intents   []types.GatewayIntent
	ws        *WebSocket
	wschannel chan webSocketPayload
	readyChan chan struct{}
	done      chan struct{}
	// guildCache map[string]Guild
}

// type guildMembersChunkEvent struct {
// 	GuildID string        `json:"guild_id"`
// 	Members []GuildMember `json:"members"`
// }

type PresenceUpdate struct {
	Since      int64      `json:"since"`
	Activities []Activity `json:"activities"`
	Status     string     `json:"status"`
	AFK        bool       `json:"afk"`
}

func (c *Client) Connect() error {
	c.ws = newWebSocket()
	c.done = make(chan struct{})
	c.readyChan = make(chan struct{})
	c.wschannel = make(chan webSocketPayload)
	go func() {
		c.ws.Connect(os.Getenv("GODISCORD_TOKEN"), c.Intents, c.wschannel)
		close(c.wschannel)
	}()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("panic recovered:", r)
			}
		}()
		for payload := range c.wschannel {
			if payload.OP == 0 {
				switch payload.EventName {
				case "READY":
					var thing map[string]any
					err := json.Unmarshal(payload.Data, &thing)
					if err != nil {
						continue
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
					close(c.readyChan)
					c.Emit("READY", c)
				case "MESSAGE_CREATE":
					var message Message
					json.Unmarshal(payload.Data, &message)
					ptr_channel, err := c.GetTextChannelByID(message.ChannelID)
					if err != nil {
						continue
					}
					if ptr_channel == nil {
						ptr_channel = &TextChannel{}
					}
					message.Channel = *ptr_channel
					ptr_owner, err := message.Channel.Guild.GetMemberByID(message.Channel.Guild.OwnerID)
					if err != nil {
						continue
					}
					if ptr_owner == nil {
						ptr_owner = &GuildMember{}
					}
					message.Channel.Guild.Owner = *ptr_owner
					message.Channel.Guild.Prunes = PruneManager{
						Guild: &message.Channel.Guild,
					}
					message.Channel.Guild.Me = localGuildMember{
						GuildId: message.Channel.GuildID,
					}
					message.Channel.Guild.Roles = RoleManager{
						GuildID: message.Channel.GuildID,
					}
					c.Emit("MESSAGE_CREATE", message, c)
				case "INTERACTION_CREATE":
					var interaction BaseInteraction
					json.Unmarshal(payload.Data, &interaction)
					fmt.Println(interaction.Type, enums.InteractionType.Message)
					c.Emit("INTERACTION_CREATE", interaction, c)
				case "MESSAGE_UPDATE":
					var message Message
					json.Unmarshal(payload.Data, &message)
					ptr_channel, err := c.GetTextChannelByID(message.ChannelID)
					if err != nil {
						continue
					}
					if ptr_channel == nil {
						ptr_channel = &TextChannel{}
					}
					message.Channel = *ptr_channel
					ptr_owner, err := message.Channel.Guild.GetMemberByID(message.Channel.Guild.OwnerID)
					if err != nil {
						continue
					}
					if ptr_owner == nil {
						ptr_owner = &GuildMember{}
					}
					message.Channel.Guild.Owner = *ptr_owner
					message.Channel.Guild.Prunes = PruneManager{
						Guild: &message.Channel.Guild,
					}
					message.Channel.Guild.Me = localGuildMember{
						GuildId: message.Channel.GuildID,
					}
					message.Channel.Guild.Roles = RoleManager{
						GuildID: message.Channel.GuildID,
					}
					c.Emit("MESSAGE_UPDATE", message, c)
				case "MESSAGE_REACTION_ADD":
					var message Message
					json.Unmarshal(payload.Data, &message)
					ptr_channel, err := c.GetTextChannelByID(message.ChannelID)
					if err != nil {
						continue
					}
					if ptr_channel == nil {
						ptr_channel = &TextChannel{}
					}
					message.Channel = *ptr_channel
					ptr_owner, err := message.Channel.Guild.GetMemberByID(message.Channel.Guild.OwnerID)
					if err != nil {
						continue
					}
					if ptr_owner == nil {
						ptr_owner = &GuildMember{}
					}
					message.Channel.Guild.Owner = *ptr_owner
					message.Channel.Guild.Prunes = PruneManager{
						Guild: &message.Channel.Guild,
					}
					message.Channel.Guild.Me = localGuildMember{
						GuildId: message.Channel.GuildID,
					}
					message.Channel.Guild.Roles = RoleManager{
						GuildID: message.Channel.GuildID,
					}
					c.Emit("MESSAGE_REACTION_ADD", message, c)
				case "GUILD_CREATE":
					var guild Guild
					json.Unmarshal(payload.Data, &guild)
					ptr_owner, err := guild.GetMemberByID(guild.OwnerID)
					if err != nil {
						continue
					}
					if ptr_owner == nil {
						ptr_owner = &GuildMember{}
					}
					guild.Owner = *ptr_owner
					guild.Prunes = PruneManager{
						Guild: &guild,
					}
					// c.guildCache[guild.ID] = guild
					guild.Me = localGuildMember{
						GuildId: guild.ID,
					}
					guild.Roles = RoleManager{
						GuildID: guild.ID,
					}
					c.Emit("GUILD_CREATE", guild, c)
				case "GUILD_DELETE":
					var guild Guild
					json.Unmarshal(payload.Data, &guild)
					ptr_owner, err := guild.GetMemberByID(guild.OwnerID)
					if err != nil {
						continue
					}
					if ptr_owner == nil {
						ptr_owner = &GuildMember{}
					}
					guild.Owner = *ptr_owner
					guild.Me = localGuildMember{
						GuildId: guild.ID,
					}
					guild.Roles = RoleManager{
						GuildID: guild.ID,
					}
					c.Emit("GUILD_DELETE", guild, c)
				case "GUILD_UPDATE":
					var guild Guild
					json.Unmarshal(payload.Data, &guild)
					ptr_owner, err := guild.GetMemberByID(guild.OwnerID)
					if err != nil {
						continue
					}
					if ptr_owner == nil {
						ptr_owner = &GuildMember{}
					}
					guild.Owner = *ptr_owner
					guild.Prunes = PruneManager{
						Guild: &guild,
					}
					guild.Me = localGuildMember{
						GuildId: guild.ID,
					}
					guild.Roles = RoleManager{
						GuildID: guild.ID,
					}
					c.Emit("GUILD_UPDATE", guild, c)
				case "GUILD_ROLE_CREATE":
					var role Role
					err := json.Unmarshal(payload.Data, &role)
					if err != nil {
						continue
					}
					c.Emit("GUILD_ROLE_CREATE", role, c)
				case "GUILD_ROLE_DELETE":
					var role Role
					err := json.Unmarshal(payload.Data, &role)
					if err != nil {
						continue
					}
					// botMember, err := role.Guild.GetMemberByID(c.User.ID)
					// if err != nil {
					//   return err
					// }
					// guild.Me = *botMember

					c.Emit("GUILD_ROLE_DELETE", role, c)
				case "GUILD_ROLE_UPDATE":
					var role Role
					err := json.Unmarshal(payload.Data, &role)
					if err != nil {
						continue
					}
					c.Emit("GUILD_ROLE_UPDATE", role, c)
				case "CHANNEL_CREATE":
					var channel Channel
					err := json.Unmarshal(payload.Data, &channel)
					if err != nil {
						continue
					}
					channel.Guild.Me = localGuildMember{
						GuildId: channel.GuildID,
					}
					channel.Guild.Prunes = PruneManager{
						Guild: &channel.Guild,
					}
					channel.Guild.Roles = RoleManager{
						GuildID: channel.GuildID,
					}
					c.Emit("CHANNEL_CREATE", channel, c)
				case "GUILD_MEMBER_ADD":
					var guildMember GuildMember
					if err := json.Unmarshal(payload.Data, &guildMember); err != nil {
						continue
					}
					if guildMember.GuildID == "" {
						continue
					}
					guildMember.RoleManager.GuildID = guildMember.GuildID
					guildMember.RoleManager.MemberID = guildMember.User.ID
					c.Emit("GUILD_MEMBER_ADD", guildMember, c)
				default:
					fmt.Println("Event:", payload.EventName)
				}
			}
		}
	}()

	<-c.done

	return nil
}

func (c Client) GetTextChannelByID(ID string) (*TextChannel, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/channels/%s", API_URL, ID), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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
	if channel.Type != types.TextChannel {
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
	request.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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
	body_reader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/guilds", API_URL), body_reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
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

func (c Client) SetPresence(Options PresenceUpdate) error {

	select {
	case <-c.readyChan:
	case <-time.After(10 * time.Second):
		return fmt.Errorf("timeout waiting for ready")
	}

	for i := range Options.Activities {
		if Options.Activities[i].CreatedAt <= 0 {
			Options.Activities[i].CreatedAt = time.Now().Unix() // <-- millisecondes ici !
		}
	}
	if Options.Since <= 0 {
		Options.Since = time.Now().Unix()
	}
	err := c.ws.SendEvent(3, Options)
	if err != nil {
		fmt.Println("Erreur SendEvent:", err)
	}
	return err
}

func (c Client) GetWSPing() int {
	return int(c.ws.Ping)
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
