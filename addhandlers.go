package dream

import (
	"github.com/Necroforger/discordgo"
)

// -------------------------------------
// AUTO GENERATED CODE. DO NOT EDIT
// --------------------------------------

//AddHandler wraps the discordgo addhandler function to return a Bot object.
func (b *Bot) AddHandler(i interface{}) {
	switch t := i.(type) {

	// ChannelCreate
	case func(*Bot, *discordgo.ChannelCreate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.ChannelCreate) {
			t(b, data)
		})

	// ChannelDelete
	case func(*Bot, *discordgo.ChannelDelete):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.ChannelDelete) {
			t(b, data)
		})

	// ChannelPinsUpdate
	case func(*Bot, *discordgo.ChannelPinsUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.ChannelPinsUpdate) {
			t(b, data)
		})

	// ChannelUpdate
	case func(*Bot, *discordgo.ChannelUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.ChannelUpdate) {
			t(b, data)
		})

	// GuildBanAdd
	case func(*Bot, *discordgo.GuildBanAdd):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildBanAdd) {
			t(b, data)
		})

	// GuildBanRemove
	case func(*Bot, *discordgo.GuildBanRemove):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildBanRemove) {
			t(b, data)
		})

	// GuildCreate
	case func(*Bot, *discordgo.GuildCreate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildCreate) {
			t(b, data)
		})

	// GuildDelete
	case func(*Bot, *discordgo.GuildDelete):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildDelete) {
			t(b, data)
		})

	// GuildEmojisUpdate
	case func(*Bot, *discordgo.GuildEmojisUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildEmojisUpdate) {
			t(b, data)
		})

	// GuildIntegrationsUpdate
	case func(*Bot, *discordgo.GuildIntegrationsUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildIntegrationsUpdate) {
			t(b, data)
		})

	// GuildMemberAdd
	case func(*Bot, *discordgo.GuildMemberAdd):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildMemberAdd) {
			t(b, data)
		})

	// GuildMemberRemove
	case func(*Bot, *discordgo.GuildMemberRemove):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildMemberRemove) {
			t(b, data)
		})

	// GuildMemberUpdate
	case func(*Bot, *discordgo.GuildMemberUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildMemberUpdate) {
			t(b, data)
		})

	// GuildMembersChunk
	case func(*Bot, *discordgo.GuildMembersChunk):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildMembersChunk) {
			t(b, data)
		})

	// GuildRoleCreate
	case func(*Bot, *discordgo.GuildRoleCreate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildRoleCreate) {
			t(b, data)
		})

	// GuildRoleDelete
	case func(*Bot, *discordgo.GuildRoleDelete):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildRoleDelete) {
			t(b, data)
		})

	// GuildRoleUpdate
	case func(*Bot, *discordgo.GuildRoleUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildRoleUpdate) {
			t(b, data)
		})

	// GuildUpdate
	case func(*Bot, *discordgo.GuildUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.GuildUpdate) {
			t(b, data)
		})

	// MessageAck
	case func(*Bot, *discordgo.MessageAck):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.MessageAck) {
			t(b, data)
		})

	// MessageCreate
	case func(*Bot, *discordgo.MessageCreate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.MessageCreate) {
			t(b, data)
		})

	// MessageDelete
	case func(*Bot, *discordgo.MessageDelete):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.MessageDelete) {
			t(b, data)
		})

	// MessageDeleteBulk
	case func(*Bot, *discordgo.MessageDeleteBulk):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.MessageDeleteBulk) {
			t(b, data)
		})

	// MessageReactionAdd
	case func(*Bot, *discordgo.MessageReactionAdd):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.MessageReactionAdd) {
			t(b, data)
		})

	// MessageReactionRemove
	case func(*Bot, *discordgo.MessageReactionRemove):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.MessageReactionRemove) {
			t(b, data)
		})

	// MessageReactionRemoveAll
	case func(*Bot, *discordgo.MessageReactionRemoveAll):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.MessageReactionRemoveAll) {
			t(b, data)
		})

	// MessageUpdate
	case func(*Bot, *discordgo.MessageUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.MessageUpdate) {
			t(b, data)
		})

	// PresenceUpdate
	case func(*Bot, *discordgo.PresenceUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.PresenceUpdate) {
			t(b, data)
		})

	// PresencesReplace
	case func(*Bot, *discordgo.PresencesReplace):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.PresencesReplace) {
			t(b, data)
		})

	// Ready
	case func(*Bot, *discordgo.Ready):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.Ready) {
			t(b, data)
		})

	// RelationshipAdd
	case func(*Bot, *discordgo.RelationshipAdd):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.RelationshipAdd) {
			t(b, data)
		})

	// RelationshipRemove
	case func(*Bot, *discordgo.RelationshipRemove):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.RelationshipRemove) {
			t(b, data)
		})

	// Resumed
	case func(*Bot, *discordgo.Resumed):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.Resumed) {
			t(b, data)
		})

	// TypingStart
	case func(*Bot, *discordgo.TypingStart):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.TypingStart) {
			t(b, data)
		})

	// UserGuildSettingsUpdate
	case func(*Bot, *discordgo.UserGuildSettingsUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.UserGuildSettingsUpdate) {
			t(b, data)
		})

	// UserNoteUpdate
	case func(*Bot, *discordgo.UserNoteUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.UserNoteUpdate) {
			t(b, data)
		})

	// UserSettingsUpdate
	case func(*Bot, *discordgo.UserSettingsUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.UserSettingsUpdate) {
			t(b, data)
		})

	// UserUpdate
	case func(*Bot, *discordgo.UserUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.UserUpdate) {
			t(b, data)
		})

	// VoiceServerUpdate
	case func(*Bot, *discordgo.VoiceServerUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.VoiceServerUpdate) {
			t(b, data)
		})

	// VoiceStateUpdate
	case func(*Bot, *discordgo.VoiceStateUpdate):
		b.DG.AddHandler(func(s *discordgo.Session, data *discordgo.VoiceStateUpdate) {
			t(b, data)
		})

	}
}

//AddHandlerOnce wraps the discordgo AddHandlerOnce function to return a Bot object
func (b *Bot) AddHandlerOnce(i interface{}) {
	switch t := i.(type) {

	// ChannelCreate
	case func(*Bot, *discordgo.ChannelCreate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.ChannelCreate) {
			t(b, data)
		})

	// ChannelDelete
	case func(*Bot, *discordgo.ChannelDelete):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.ChannelDelete) {
			t(b, data)
		})

	// ChannelPinsUpdate
	case func(*Bot, *discordgo.ChannelPinsUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.ChannelPinsUpdate) {
			t(b, data)
		})

	// ChannelUpdate
	case func(*Bot, *discordgo.ChannelUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.ChannelUpdate) {
			t(b, data)
		})

	// GuildBanAdd
	case func(*Bot, *discordgo.GuildBanAdd):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildBanAdd) {
			t(b, data)
		})

	// GuildBanRemove
	case func(*Bot, *discordgo.GuildBanRemove):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildBanRemove) {
			t(b, data)
		})

	// GuildCreate
	case func(*Bot, *discordgo.GuildCreate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildCreate) {
			t(b, data)
		})

	// GuildDelete
	case func(*Bot, *discordgo.GuildDelete):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildDelete) {
			t(b, data)
		})

	// GuildEmojisUpdate
	case func(*Bot, *discordgo.GuildEmojisUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildEmojisUpdate) {
			t(b, data)
		})

	// GuildIntegrationsUpdate
	case func(*Bot, *discordgo.GuildIntegrationsUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildIntegrationsUpdate) {
			t(b, data)
		})

	// GuildMemberAdd
	case func(*Bot, *discordgo.GuildMemberAdd):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildMemberAdd) {
			t(b, data)
		})

	// GuildMemberRemove
	case func(*Bot, *discordgo.GuildMemberRemove):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildMemberRemove) {
			t(b, data)
		})

	// GuildMemberUpdate
	case func(*Bot, *discordgo.GuildMemberUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildMemberUpdate) {
			t(b, data)
		})

	// GuildMembersChunk
	case func(*Bot, *discordgo.GuildMembersChunk):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildMembersChunk) {
			t(b, data)
		})

	// GuildRoleCreate
	case func(*Bot, *discordgo.GuildRoleCreate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildRoleCreate) {
			t(b, data)
		})

	// GuildRoleDelete
	case func(*Bot, *discordgo.GuildRoleDelete):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildRoleDelete) {
			t(b, data)
		})

	// GuildRoleUpdate
	case func(*Bot, *discordgo.GuildRoleUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildRoleUpdate) {
			t(b, data)
		})

	// GuildUpdate
	case func(*Bot, *discordgo.GuildUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.GuildUpdate) {
			t(b, data)
		})

	// MessageAck
	case func(*Bot, *discordgo.MessageAck):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.MessageAck) {
			t(b, data)
		})

	// MessageCreate
	case func(*Bot, *discordgo.MessageCreate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.MessageCreate) {
			t(b, data)
		})

	// MessageDelete
	case func(*Bot, *discordgo.MessageDelete):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.MessageDelete) {
			t(b, data)
		})

	// MessageDeleteBulk
	case func(*Bot, *discordgo.MessageDeleteBulk):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.MessageDeleteBulk) {
			t(b, data)
		})

	// MessageReactionAdd
	case func(*Bot, *discordgo.MessageReactionAdd):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.MessageReactionAdd) {
			t(b, data)
		})

	// MessageReactionRemove
	case func(*Bot, *discordgo.MessageReactionRemove):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.MessageReactionRemove) {
			t(b, data)
		})

	// MessageReactionRemoveAll
	case func(*Bot, *discordgo.MessageReactionRemoveAll):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.MessageReactionRemoveAll) {
			t(b, data)
		})

	// MessageUpdate
	case func(*Bot, *discordgo.MessageUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.MessageUpdate) {
			t(b, data)
		})

	// PresenceUpdate
	case func(*Bot, *discordgo.PresenceUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.PresenceUpdate) {
			t(b, data)
		})

	// PresencesReplace
	case func(*Bot, *discordgo.PresencesReplace):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.PresencesReplace) {
			t(b, data)
		})

	// Ready
	case func(*Bot, *discordgo.Ready):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.Ready) {
			t(b, data)
		})

	// RelationshipAdd
	case func(*Bot, *discordgo.RelationshipAdd):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.RelationshipAdd) {
			t(b, data)
		})

	// RelationshipRemove
	case func(*Bot, *discordgo.RelationshipRemove):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.RelationshipRemove) {
			t(b, data)
		})

	// Resumed
	case func(*Bot, *discordgo.Resumed):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.Resumed) {
			t(b, data)
		})

	// TypingStart
	case func(*Bot, *discordgo.TypingStart):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.TypingStart) {
			t(b, data)
		})

	// UserGuildSettingsUpdate
	case func(*Bot, *discordgo.UserGuildSettingsUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.UserGuildSettingsUpdate) {
			t(b, data)
		})

	// UserNoteUpdate
	case func(*Bot, *discordgo.UserNoteUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.UserNoteUpdate) {
			t(b, data)
		})

	// UserSettingsUpdate
	case func(*Bot, *discordgo.UserSettingsUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.UserSettingsUpdate) {
			t(b, data)
		})

	// UserUpdate
	case func(*Bot, *discordgo.UserUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.UserUpdate) {
			t(b, data)
		})

	// VoiceServerUpdate
	case func(*Bot, *discordgo.VoiceServerUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.VoiceServerUpdate) {
			t(b, data)
		})

	// VoiceStateUpdate
	case func(*Bot, *discordgo.VoiceStateUpdate):
		b.DG.AddHandlerOnce(func(s *discordgo.Session, data *discordgo.VoiceStateUpdate) {
			t(b, data)
		})

	}
}
