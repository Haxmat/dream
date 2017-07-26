package dream

import (
	"github.com/Necroforger/discordgo"
)

// -------------------------------------
// AUTO GENERATED CODE. DO NOT EDIT
// --------------------------------------

//AddHandler wraps the discordgo addhandler function to return a Session object.
func (s *Session) AddHandler(i interface{}) {
	switch t := i.(type) {

	// ChannelCreate
	case func(*Session, *discordgo.ChannelCreate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.ChannelCreate) {
			t(s, data)
		})

	// ChannelDelete
	case func(*Session, *discordgo.ChannelDelete):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.ChannelDelete) {
			t(s, data)
		})

	// ChannelPinsUpdate
	case func(*Session, *discordgo.ChannelPinsUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.ChannelPinsUpdate) {
			t(s, data)
		})

	// ChannelUpdate
	case func(*Session, *discordgo.ChannelUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.ChannelUpdate) {
			t(s, data)
		})

	// GuildBanAdd
	case func(*Session, *discordgo.GuildBanAdd):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildBanAdd) {
			t(s, data)
		})

	// GuildBanRemove
	case func(*Session, *discordgo.GuildBanRemove):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildBanRemove) {
			t(s, data)
		})

	// GuildCreate
	case func(*Session, *discordgo.GuildCreate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildCreate) {
			t(s, data)
		})

	// GuildDelete
	case func(*Session, *discordgo.GuildDelete):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildDelete) {
			t(s, data)
		})

	// GuildEmojisUpdate
	case func(*Session, *discordgo.GuildEmojisUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildEmojisUpdate) {
			t(s, data)
		})

	// GuildIntegrationsUpdate
	case func(*Session, *discordgo.GuildIntegrationsUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildIntegrationsUpdate) {
			t(s, data)
		})

	// GuildMemberAdd
	case func(*Session, *discordgo.GuildMemberAdd):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildMemberAdd) {
			t(s, data)
		})

	// GuildMemberRemove
	case func(*Session, *discordgo.GuildMemberRemove):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildMemberRemove) {
			t(s, data)
		})

	// GuildMemberUpdate
	case func(*Session, *discordgo.GuildMemberUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildMemberUpdate) {
			t(s, data)
		})

	// GuildMembersChunk
	case func(*Session, *discordgo.GuildMembersChunk):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildMembersChunk) {
			t(s, data)
		})

	// GuildRoleCreate
	case func(*Session, *discordgo.GuildRoleCreate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildRoleCreate) {
			t(s, data)
		})

	// GuildRoleDelete
	case func(*Session, *discordgo.GuildRoleDelete):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildRoleDelete) {
			t(s, data)
		})

	// GuildRoleUpdate
	case func(*Session, *discordgo.GuildRoleUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildRoleUpdate) {
			t(s, data)
		})

	// GuildUpdate
	case func(*Session, *discordgo.GuildUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.GuildUpdate) {
			t(s, data)
		})

	// MessageAck
	case func(*Session, *discordgo.MessageAck):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.MessageAck) {
			t(s, data)
		})

	// MessageCreate
	case func(*Session, *discordgo.MessageCreate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.MessageCreate) {
			t(s, data)
		})

	// MessageDelete
	case func(*Session, *discordgo.MessageDelete):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.MessageDelete) {
			t(s, data)
		})

	// MessageDeleteBulk
	case func(*Session, *discordgo.MessageDeleteBulk):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.MessageDeleteBulk) {
			t(s, data)
		})

	// MessageReactionAdd
	case func(*Session, *discordgo.MessageReactionAdd):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.MessageReactionAdd) {
			t(s, data)
		})

	// MessageReactionRemove
	case func(*Session, *discordgo.MessageReactionRemove):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.MessageReactionRemove) {
			t(s, data)
		})

	// MessageReactionRemoveAll
	case func(*Session, *discordgo.MessageReactionRemoveAll):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.MessageReactionRemoveAll) {
			t(s, data)
		})

	// MessageUpdate
	case func(*Session, *discordgo.MessageUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.MessageUpdate) {
			t(s, data)
		})

	// PresenceUpdate
	case func(*Session, *discordgo.PresenceUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.PresenceUpdate) {
			t(s, data)
		})

	// PresencesReplace
	case func(*Session, *discordgo.PresencesReplace):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.PresencesReplace) {
			t(s, data)
		})

	// Ready
	case func(*Session, *discordgo.Ready):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.Ready) {
			t(s, data)
		})

	// RelationshipAdd
	case func(*Session, *discordgo.RelationshipAdd):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.RelationshipAdd) {
			t(s, data)
		})

	// RelationshipRemove
	case func(*Session, *discordgo.RelationshipRemove):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.RelationshipRemove) {
			t(s, data)
		})

	// Resumed
	case func(*Session, *discordgo.Resumed):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.Resumed) {
			t(s, data)
		})

	// TypingStart
	case func(*Session, *discordgo.TypingStart):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.TypingStart) {
			t(s, data)
		})

	// UserGuildSettingsUpdate
	case func(*Session, *discordgo.UserGuildSettingsUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.UserGuildSettingsUpdate) {
			t(s, data)
		})

	// UserNoteUpdate
	case func(*Session, *discordgo.UserNoteUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.UserNoteUpdate) {
			t(s, data)
		})

	// UserSettingsUpdate
	case func(*Session, *discordgo.UserSettingsUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.UserSettingsUpdate) {
			t(s, data)
		})

	// UserUpdate
	case func(*Session, *discordgo.UserUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.UserUpdate) {
			t(s, data)
		})

	// VoiceServerUpdate
	case func(*Session, *discordgo.VoiceServerUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.VoiceServerUpdate) {
			t(s, data)
		})

	// VoiceStateUpdate
	case func(*Session, *discordgo.VoiceStateUpdate):
		s.DG.AddHandler(func(_ *discordgo.Session, data *discordgo.VoiceStateUpdate) {
			t(s, data)
		})

	}
}

//AddHandlerOnce wraps the discordgo AddHandlerOnce function to return a Session object
func (s *Session) AddHandlerOnce(i interface{}) {
	switch t := i.(type) {

	// ChannelCreate
	case func(*Session, *discordgo.ChannelCreate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.ChannelCreate) {
			t(s, data)
		})

	// ChannelDelete
	case func(*Session, *discordgo.ChannelDelete):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.ChannelDelete) {
			t(s, data)
		})

	// ChannelPinsUpdate
	case func(*Session, *discordgo.ChannelPinsUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.ChannelPinsUpdate) {
			t(s, data)
		})

	// ChannelUpdate
	case func(*Session, *discordgo.ChannelUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.ChannelUpdate) {
			t(s, data)
		})

	// GuildBanAdd
	case func(*Session, *discordgo.GuildBanAdd):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildBanAdd) {
			t(s, data)
		})

	// GuildBanRemove
	case func(*Session, *discordgo.GuildBanRemove):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildBanRemove) {
			t(s, data)
		})

	// GuildCreate
	case func(*Session, *discordgo.GuildCreate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildCreate) {
			t(s, data)
		})

	// GuildDelete
	case func(*Session, *discordgo.GuildDelete):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildDelete) {
			t(s, data)
		})

	// GuildEmojisUpdate
	case func(*Session, *discordgo.GuildEmojisUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildEmojisUpdate) {
			t(s, data)
		})

	// GuildIntegrationsUpdate
	case func(*Session, *discordgo.GuildIntegrationsUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildIntegrationsUpdate) {
			t(s, data)
		})

	// GuildMemberAdd
	case func(*Session, *discordgo.GuildMemberAdd):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildMemberAdd) {
			t(s, data)
		})

	// GuildMemberRemove
	case func(*Session, *discordgo.GuildMemberRemove):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildMemberRemove) {
			t(s, data)
		})

	// GuildMemberUpdate
	case func(*Session, *discordgo.GuildMemberUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildMemberUpdate) {
			t(s, data)
		})

	// GuildMembersChunk
	case func(*Session, *discordgo.GuildMembersChunk):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildMembersChunk) {
			t(s, data)
		})

	// GuildRoleCreate
	case func(*Session, *discordgo.GuildRoleCreate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildRoleCreate) {
			t(s, data)
		})

	// GuildRoleDelete
	case func(*Session, *discordgo.GuildRoleDelete):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildRoleDelete) {
			t(s, data)
		})

	// GuildRoleUpdate
	case func(*Session, *discordgo.GuildRoleUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildRoleUpdate) {
			t(s, data)
		})

	// GuildUpdate
	case func(*Session, *discordgo.GuildUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.GuildUpdate) {
			t(s, data)
		})

	// MessageAck
	case func(*Session, *discordgo.MessageAck):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.MessageAck) {
			t(s, data)
		})

	// MessageCreate
	case func(*Session, *discordgo.MessageCreate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.MessageCreate) {
			t(s, data)
		})

	// MessageDelete
	case func(*Session, *discordgo.MessageDelete):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.MessageDelete) {
			t(s, data)
		})

	// MessageDeleteBulk
	case func(*Session, *discordgo.MessageDeleteBulk):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.MessageDeleteBulk) {
			t(s, data)
		})

	// MessageReactionAdd
	case func(*Session, *discordgo.MessageReactionAdd):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.MessageReactionAdd) {
			t(s, data)
		})

	// MessageReactionRemove
	case func(*Session, *discordgo.MessageReactionRemove):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.MessageReactionRemove) {
			t(s, data)
		})

	// MessageReactionRemoveAll
	case func(*Session, *discordgo.MessageReactionRemoveAll):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.MessageReactionRemoveAll) {
			t(s, data)
		})

	// MessageUpdate
	case func(*Session, *discordgo.MessageUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.MessageUpdate) {
			t(s, data)
		})

	// PresenceUpdate
	case func(*Session, *discordgo.PresenceUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.PresenceUpdate) {
			t(s, data)
		})

	// PresencesReplace
	case func(*Session, *discordgo.PresencesReplace):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.PresencesReplace) {
			t(s, data)
		})

	// Ready
	case func(*Session, *discordgo.Ready):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.Ready) {
			t(s, data)
		})

	// RelationshipAdd
	case func(*Session, *discordgo.RelationshipAdd):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.RelationshipAdd) {
			t(s, data)
		})

	// RelationshipRemove
	case func(*Session, *discordgo.RelationshipRemove):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.RelationshipRemove) {
			t(s, data)
		})

	// Resumed
	case func(*Session, *discordgo.Resumed):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.Resumed) {
			t(s, data)
		})

	// TypingStart
	case func(*Session, *discordgo.TypingStart):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.TypingStart) {
			t(s, data)
		})

	// UserGuildSettingsUpdate
	case func(*Session, *discordgo.UserGuildSettingsUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.UserGuildSettingsUpdate) {
			t(s, data)
		})

	// UserNoteUpdate
	case func(*Session, *discordgo.UserNoteUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.UserNoteUpdate) {
			t(s, data)
		})

	// UserSettingsUpdate
	case func(*Session, *discordgo.UserSettingsUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.UserSettingsUpdate) {
			t(s, data)
		})

	// UserUpdate
	case func(*Session, *discordgo.UserUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.UserUpdate) {
			t(s, data)
		})

	// VoiceServerUpdate
	case func(*Session, *discordgo.VoiceServerUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.VoiceServerUpdate) {
			t(s, data)
		})

	// VoiceStateUpdate
	case func(*Session, *discordgo.VoiceStateUpdate):
		s.DG.AddHandlerOnce(func(_ *discordgo.Session, data *discordgo.VoiceStateUpdate) {
			t(s, data)
		})

	}
}
