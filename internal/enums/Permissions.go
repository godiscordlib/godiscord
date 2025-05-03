package enums

import (
	"encoding/json"
	"strconv"
)

type Permission int

func (p Permission) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(int64(p), 10))
}
func (p *Permission) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*p = Permission(v)
	return nil
}

const (
	Permission_CreateInstantInvite Permission = 1 << iota
	Permission_KickMembers
	Permission_BanMembers
	Permission_Administrator
	Permission_ManageChannels
	Permission_ManageGuild
	Permission_AddReactions
	Permission_ViewAuditLog
	Permission_PrioritySpeaker
	Permission_Stream
	Permission_ViewChannel
	Permission_SendMessages
	Permission_SendTTSMessages
	Permission_ManageMessages
	Permission_LinkEmbeds
	Permission_AttachFiles
	Permission_ReadMessageHistory
	Permission_MentionEveryone
	Permission_UseExternalEmojis
	Permission_ViewGuildInsights
	Permission_Connect
	Permission_Speak
	Permission_MuteMembers
	Permission_DeafenMembers
	Permission_MoveMembers
	Permission_UseVoiceActivityDetection
	Permission_ChangeNickname
	Permission_ManageNicknames
	Permission_ManageRoles
	Permission_ManageWebhooks
	Permission_ManageGuildExpressions
	Permission_UseApplicationCommands
	Permission_RequestToSpeak
	Permission_ManageEvents
	Permission_ManageThreads
	Permission_CreatePublicThread
	Permission_CreatePrivateThread
	Permission_UseExternalStickers
	Permission_SendMessagesInThreads
	Permission_UseEmbeddedActivities
	Permission_ModerateMembers
	Permission_ViewMonetizationAnalytics
	Permission_UseSoundboard
	Permission_CreateGuildExpressions
	Permission_CreateEvents
	Permission_UseExternalSounds
	Permission_SendVoiceMessages
	Permission_SendPolls
	Permission_UseExternalApps
)
