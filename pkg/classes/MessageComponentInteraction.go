package classes

import "github.com/godiscordlib/godiscord/pkg/types"

// "encoding/json"

type MessageComponentInteraction struct {
	ID   string              `json:"custom_id"`
	Type types.ComponentType `json:"component_type"`
}

type resolvedData struct {
	Users       map[string]User        `json:"users"`
	Members     map[string]GuildMember `json:"members"`
	Roles       map[string]Role        `json:"roles"`
	Channels    map[string]Channel     `json:"channels"`
	Messages    map[string]Message     `json:"messages"`
	Attachments map[string]Attachment  `json:"attachments"`
}
