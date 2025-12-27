package classes

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/godiscordlib/godiscord/pkg/enums"
	"github.com/godiscordlib/godiscord/pkg/types"
)

type BaseInteraction struct {
	Type      types.InteractionResponseType `json:"type"`
	Token     string                        `json:"token"`
	Member    GuildMember                   `json:"member"`
	ID        string                        `json:"id"`
	GuildID   string                        `json:"guild_id"`
	Data      baseInteractionData           `json:"data"`
	ChannelID string                        `json:"channel_id"`
	Channel   ChannelInt
	Guild     Guild
}

type baseInteractionData struct {
	Type          int                    `json:"type"`
	Name          *string                `json:"name"`
	ID            *string                `json:"id"`
	CustomID      *string                `json:"custom_id"`
	Values        *[]string              `json:"values"`
	Resolved      *resolvedData          `json:"resolved"`
	ComponentType *types.ComponentType   `json:"component_type"`
	Value         *any                   `json:"value"`
	Focused       *bool                  `json:"focused"`
	Options       *[]baseInteractionData `json:"options"`
}

type BaseComponent interface {
	GetType() types.ComponentType
}

func (bi BaseInteraction) GetName() string {
	if bi.Type != enums.InteractionResponseType.ApplicationCommand {
		return ""
	}

	return *bi.Data.Name
}

func (bi BaseInteraction) Values() []string {
	if bi.Type != enums.InteractionResponseType.MessageComponent {
		return []string{}
	}

	return *bi.Data.Values
}

func (bi BaseInteraction) Resolved() resolvedData {
	if bi.Type != enums.InteractionResponseType.MessageComponent {
		return *new(resolvedData)
	}

	return *bi.Data.Resolved
}

func (bi BaseInteraction) GetUser(OptionName string) *User {
	if len(*bi.Data.Options) == 0 {
		return nil
	}
	for _, o := range *bi.Data.Options {
		if *o.Name == OptionName {
			optionValue := *o.Value
			req, err := http.NewRequest("GET", API_URL+"/users/"+optionValue.(string), nil)
			if err != nil {
				return nil
			}
			req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return nil
			}
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return nil
			}
			var user User
			if err = json.Unmarshal(body, &user); err != nil {
				return nil
			}
			return &user
		}
	}

	return nil
}

func (bi BaseInteraction) GetInteger(OptionName string) *int {
	if len(*bi.Data.Options) == 0 {
		return nil
	}
	for _, o := range *bi.Data.Options {
		if *o.Name == OptionName {
			optionValue := *o.Value
			return optionValue.(*int)
		}
	}
	return nil
}
