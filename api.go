package dream

import (
	"errors"
	"fmt"
	"log"
	"sort"

	"github.com/jonas747/dca"

	//Blank imports included for decoding a user's avatar into an image.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"io"
	"os"

	"github.com/bwmarrin/discordgo"
)

/*                   TODO

TODO: Improve GuildMemberRoleAdd, GuildMemberRoleRemove methods to accept interfaces to be more convenient.
	  It could possibly obtain the guildID from an interface.

*/

//Error values//
var (
	ErrInvalidType      = errors.New("err: invalid type")
	ErrNotFound         = errors.New("err: not found")
	ErrInvalidArgLength = errors.New("err: invalid argument length")
)

// MessageFromInterface returns the underlying message struct of an event relating to messages
// Such as MessageCreate, MessageDelete, MessageUpdate.
func (s *Session) MessageFromInterface(i interface{}) (*discordgo.Message, error) {
	switch t := i.(type) {
	case *discordgo.MessageCreate:
		return t.Message, nil
	case *discordgo.MessageUpdate:
		return t.Message, nil
	case *discordgo.MessageDelete:
		return t.Message, nil
	case *Message:
		return t.Message, nil
	case *discordgo.Message:
		return t, nil
	default:
		return nil, ErrInvalidType
	}
}

// ChannelID Returns the channelID from a variety of objects
func (s *Session) ChannelID(i interface{}) (string, error) {

	// Attempt to retrieve a message object from the interface
	if m, err := s.MessageFromInterface(i); err == nil {
		return m.ChannelID, nil
	}

	// If no message object is found check for other types containing a channelID.
	switch t := i.(type) {
	case string:
		return t, nil
	case *discordgo.VoiceState:
		return t.ChannelID, nil
	case *discordgo.VoiceConnection:
		return t.ChannelID, nil
	}

	return "", ErrInvalidType
}

// GuildID returns the GuildID from a variety of objects
func (s *Session) GuildID(i interface{}) (string, error) {

	// Attempt to get the Message object from the interface.
	// If it fails, check the other possible types.
	if t, err := s.MessageFromInterface(i); err == nil {

		c, err := s.Channel(t.ChannelID)
		if err != nil {
			return "", err
		}
		return c.GuildID, nil

	}

	// Check for other types
	switch t := i.(type) {
	case string:
		return t, nil

	case *discordgo.Channel:
		return t.GuildID, nil

	case *discordgo.Guild:
		return t.ID, nil

	case *discordgo.VoiceConnection:
		return t.GuildID, nil

	case *discordgo.Member:
		return t.GuildID, nil

	default:
		return "", ErrInvalidType
	}

}

// UserID returns the userID from a variety of objects.
func (s *Session) UserID(i interface{}) (userid string, err error) {

	if t, err := s.MessageFromInterface(i); err == nil {
		if t.Author != nil {
			return t.Author.ID, nil
		}
	}

	switch t := i.(type) {
	case string:
		return t, nil

	case *discordgo.User:
		return t.ID, nil

	case *discordgo.Member:
		return t.User.ID, nil

	case *discordgo.Presence:
		return t.User.ID, nil

	default:
		return "", ErrInvalidType
	}
}

// Channel is a convenience method for retrieving a channel from a variety of objects
func (s *Session) Channel(i interface{}) (*discordgo.Channel, error) {

	// Return channel if `i` is already of type channel
	if c, ok := i.(*discordgo.Channel); ok {
		return c, nil
	}

	channelid, err := s.ChannelID(i)
	if err != nil {
		return nil, err
	}
	c, err := s.DG.State.Channel(channelid)
	if err != nil {
		c, err = s.DG.Channel(channelid)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// ChannelVoiceJoin joins the specified voice channel
func (s *Session) ChannelVoiceJoin(guildID, channelID string, mute, deaf bool) (*discordgo.VoiceConnection, error) {
	vc, err := s.DG.ChannelVoiceJoin(guildID, channelID, mute, deaf)
	if err != nil {
		return nil, err
	}
	return vc, nil
}

// UserVoiceState finds a user's voice state from all the guilds in the session
func (s *Session) UserVoiceState(userID string) (*discordgo.VoiceState, error) {

	for _, v := range s.DG.State.Guilds {
		for _, c := range v.VoiceStates {
			if c.UserID == userID {
				return c, nil
			}
		}
	}

	return nil, errors.New("VoiceState not found")
}

// UserVoiceStateJoin joins a user's voice state channel.
func (s *Session) UserVoiceStateJoin(userID interface{}, mute, deaf bool) (*discordgo.VoiceConnection, error) {
	uid, err := s.UserID(userID)
	if err != nil {
		return nil, err
	}

	vs, err := s.UserVoiceState(uid)
	if err != nil {
		return nil, err
	}

	return s.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, mute, deaf)
}

func (s *Session) convertToOpus2(dst io.Writer, src io.Reader) error {
	encodingSession, err := dca.EncodeMem(src, &dca.EncodeOptions{
		Volume:           256,
		Channels:         2,
		FrameRate:        48000,
		FrameDuration:    20,
		BufferedFrames:   100,
		Bitrate:          96,
		Application:      dca.AudioApplicationAudio,
		CompressionLevel: 10,
		PacketLoss:       0,
		VBR:              true,
		RawOutput:        true,
	})
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, encodingSession)
	if err != nil {
		return err
	}

	return nil
}

// GuildAudioDispatcherStop a guild's currently playing audio dispatchers
func (s *Session) GuildAudioDispatcherStop(i interface{}) error {
	guildID, err := s.GuildID(i)
	if err != nil {
		return err
	}

	disp, err := s.audioDispatcher(guildID)
	if err != nil {
		return err
	}

	disp.Stop()

	return nil
}

// GuildAudioDispatcherPause pauses the guild's currently playing audio dispatcher
func (s *Session) GuildAudioDispatcherPause(i interface{}) error {
	guildID, err := s.GuildID(i)
	if err != nil {
		return err
	}

	disp, err := s.audioDispatcher(guildID)
	if err != nil {
		return err
	}

	disp.Pause()
	return nil
}

// GuildAudioDispatcherResume resumes the guild's currently playing audio dispatcher
func (s *Session) GuildAudioDispatcherResume(i interface{}) error {
	guildID, err := s.GuildID(i)
	if err != nil {
		return err
	}

	disp, err := s.audioDispatcher(guildID)
	if err != nil {
		return err
	}

	disp.Resume()
	return nil
}

// GuildAudioDispatcher returns the specified guild's audio dispatcher
func (s *Session) GuildAudioDispatcher(i interface{}) (*AudioDispatcher, error) {
	guildID, err := s.GuildID(i)
	if err != nil {
		return nil, err
	}

	return s.audioDispatcher(guildID)
}

// PlayStream plays an audio stream from the given io reader and uses ffmpeg to convert to a suitable format
func (s *Session) PlayStream(vc *discordgo.VoiceConnection, rd io.Reader) *AudioDispatcher {
	opusStream, wr := io.Pipe()
	go func() {
		err := s.convertToOpus2(wr, rd)
		if err != nil {
			log.Println("error converting audio to opus: ", err)
		}
		wr.Close()
	}()

	disp := NewAudioDispatcher(vc, opusStream)
	s.GuildAudioDispatcherStop(vc.GuildID)
	s.addAudioDispatcher(disp)

	go func() {
		disp.Start()
	}()

	return disp
}

// PlayRawStream plays a direct stream
func (s *Session) PlayRawStream(vc *discordgo.VoiceConnection, rd io.Reader) *AudioDispatcher {
	disp := NewAudioDispatcher(vc, rd)
	s.GuildAudioDispatcherStop(vc.GuildID)
	s.addAudioDispatcher(disp)

	go func() {
		disp.Start()
		//b.removeAudioDispatcher(disp.GuildID)
	}()

	return disp
}

// PlayFile opens a file and calls PlayStream on it
func (s *Session) PlayFile(vc *discordgo.VoiceConnection, path string) (*AudioDispatcher, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		return nil, err
	}

	disp := s.PlayStream(vc, f)
	return disp, nil
}

// PlayRawFile opens a file and calls PlayRawstream on it
func (s *Session) PlayRawFile(vc *discordgo.VoiceConnection, path string) (*AudioDispatcher, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		return nil, err
	}

	disp := s.PlayRawStream(vc, f)
	return disp, nil
}

// Guild is a convenience method for retrieving a channel from a variety of objects
func (s *Session) Guild(i interface{}) (*discordgo.Guild, error) {
	if g, ok := i.(*discordgo.Guild); ok {
		return g, nil
	}

	guildid, err := s.GuildID(i)
	if err != nil {
		return nil, err
	}
	guild, err := s.DG.State.Guild(guildid)
	if err != nil {
		guild, err = s.DG.Guild(guildid)
		if err != nil {
			return nil, err
		}
	}
	return guild, nil
}

// TODO Rework GuildPresence

// GuildPresence attempts to first find a guildMember object from the supplied arguments. If a member is found,
// It uses the members guildID and userID.
func (s *Session) GuildPresence(i ...interface{}) (*discordgo.Presence, error) {
	var userID string
	var guildID string

	if len(i) == 0 {
		return nil, ErrInvalidArgLength
	}

	// Return if the first argument is already a presence type
	if p, ok := i[0].(*discordgo.Presence); ok {
		return p, nil
	}

	//If there is only one argument, obtain both the GuildID and the UserID from it
	//Otherwise, Get the guild ID from the first argument and the UserID from the second.
	if len(i) == 1 {

		gid, err := s.GuildID(i[0])
		if err != nil {
			return nil, err
		}
		uid, err := s.UserID(i[0])
		if err != nil {
			return nil, err
		}

		userID, guildID = uid, gid

	} else {

		//Get guildID from first argument
		gid, err := s.GuildID(i[0])
		if err != nil {
			return nil, err
		}

		//Get userID from second argument
		uid, err := s.UserID(i[1])
		if err != nil {
			return nil, err
		}

		userID, guildID = uid, gid
	}

	p, err := s.DG.State.Presence(guildID, userID)
	if err == nil {
		return p, nil
	}

	return nil, ErrNotFound
}

// GuildVoiceConnection returns the voice connection object for the given guild
func (s *Session) GuildVoiceConnection(i interface{}) (*discordgo.VoiceConnection, error) {
	guildID, err := s.GuildID(i)
	if err != nil {
		return nil, err
	}

	if vc, ok := s.DG.VoiceConnections[guildID]; ok {
		return vc, nil
	}

	return nil, errors.New("Voice connection not found")
}

// GuildVoiceConnectionDisconnect finds the current guild voice connection and disconnects from it
func (s *Session) GuildVoiceConnectionDisconnect(guildID interface{}) error {
	vc, err := s.GuildVoiceConnection(guildID)
	if err != nil {
		return err
	}

	return vc.Disconnect()
}

// GuildMember is a convenience method for fetching the member object from various sources
func (s *Session) GuildMember(i ...interface{}) (*discordgo.Member, error) {
	if len(i) == 0 {
		return nil, ErrInvalidArgLength
	}

	var (
		userid  string
		guildid string
		err     error
	)

	// Return if the type is already a member object
	if v, ok := i[0].(*discordgo.Member); ok {
		return v, nil
	}

	if len(i) == 1 {
		// Attempt to get both the userID and the guildid
		// From the first argument
		guildid, err = s.GuildID(i[0])
		if err != nil {
			return nil, err
		}
		userid, err = s.UserID(i[0])
		if err != nil {
			return nil, err
		}
	} else {
		// If there are more than one argument present,
		// Attempt to get the guildID from the first one
		// And the userID from the second

		guildid, err = s.GuildID(i[0])
		if err != nil {
			return nil, err
		}

		userid, err = s.UserID(i[1])
		if err != nil {
			return nil, err
		}
	}

	// Fetch member from state and fall back to the restAPI if that fails.
	gm, err := s.DG.State.Member(guildid, userid)
	if err != nil {
		gm, err = s.DG.GuildMember(guildid, userid)
		if err != nil {
			return nil, err
		}
	}
	return gm, nil
}

// GuildRoles requests a guilds roles from the state or the API if not available in state
func (s *Session) GuildRoles(i interface{}) ([]*discordgo.Role, error) {
	guild, err := s.Guild(i)
	if err != nil {
		return nil, err
	}

	// Make sure to make a copy of guild.Roles as to not effect the underlying
	// Array in guild.Roles.
	roles := make([]*discordgo.Role, len(guild.Roles))
	copy(roles, guild.Roles)

	sort.Sort(Roles(roles))

	return roles, nil
}

// GuildMemberRoles returns an array of the roles of a member sorted by position
func (s *Session) GuildMemberRoles(i ...interface{}) ([]*discordgo.Role, error) {
	var roles []*discordgo.Role

	member, err := s.GuildMember(i...)
	if err != nil {
		return nil, err
	}

	guildRoles, err := s.GuildRoles(member.GuildID)
	if err != nil {
		return nil, err
	}

	// Find the Member's roles in the guild
	for _, gr := range guildRoles {
		for _, mr := range member.Roles {
			if gr.ID == mr {
				roles = append(roles, gr)
			}
		}
	}

	return roles, nil
}

// GuildMemberRoleAdd adds the specified role to a given member
//  guildID   : The ID of a Guild.
//  userID    : The ID of a User.
//  roleID 	  : The ID of a Role to be assigned to the user.
func (s *Session) GuildMemberRoleAdd(guildID, userID interface{}, roleID string) error {
	gid, err := s.GuildID(guildID)
	if err != nil {
		return err
	}

	uid, err := s.UserID(userID)
	if err != nil {
		return err
	}

	return s.DG.GuildMemberRoleAdd(gid, uid, roleID)
}

// GuildMemberRoleAddByName adds a role to the specified guild member
func (s *Session) GuildMemberRoleAddByName(guildID, userID interface{}, name string) error {
	guild, err := s.Guild(guildID)
	if err != nil {
		return err
	}

	uid, err := s.UserID(userID)
	if err != nil {
		return err
	}

	for _, v := range guild.Roles {
		if v.Name == name {
			return s.DG.GuildMemberRoleAdd(guild.ID, uid, v.ID)
		}
	}

	return ErrNotFound
}

// GuildMemberRoleAddByNames ...
func (s *Session) GuildMemberRoleAddByNames(guildID, userID interface{}, names ...string) (err error) {
	guild, err := s.Guild(guildID)
	if err != nil {
		return
	}

	uid, err := s.UserID(userID)
	if err != nil {
		return
	}

	for _, name := range names {
		err = s.GuildMemberRoleAddByName(guild, uid, name)
	}
	return
}

// GuildMemberRoleRemove removes the specified role to a given member
//  guildID   : The ID of a Guild.
//  userID    : The ID of a User.
//  roleID 	  : The ID of a Role to be removed from the user.
func (s *Session) GuildMemberRoleRemove(guildID, userID interface{}, roleID string) error {
	gid, err := s.GuildID(guildID)
	if err != nil {
		return err
	}

	uid, err := s.UserID(userID)
	if err != nil {
		return err
	}

	return s.DG.GuildMemberRoleRemove(gid, uid, roleID)
}

// GuildMemberRoleRemoveByName removes a role from a member by name
func (s *Session) GuildMemberRoleRemoveByName(guildID, userID interface{}, rolename string) error {
	gid, err := s.GuildID(guildID)
	if err != nil {
		return err
	}

	uid, err := s.UserID(userID)
	if err != nil {
		return err
	}

	memberRoles, err := s.GuildMemberRoles(gid, uid)
	if err != nil {
		return err
	}

	for _, v := range memberRoles {
		if v.Name == rolename {
			return s.GuildMemberRoleRemove(gid, uid, v.ID)
		}
	}

	return ErrNotFound
}

// GuildMemberRolesRemoveByName removes a list of roles by name from a guild member
func (s *Session) GuildMemberRolesRemoveByName(guildID, userID interface{}, rolenames ...string) (err error) {
	gid, err := s.GuildID(guildID)
	if err != nil {
		return err
	}

	uid, err := s.UserID(userID)
	if err != nil {
		return err
	}

	for _, rolename := range rolenames {
		err = s.GuildMemberRoleRemoveByName(gid, uid, rolename)
	}
	return
}

// RoleSettings is an object passes to GuildRoleCreate or GuildRoleEdit to
// Deal with setting the values of the command easier
// guildID   : The ID of a Guild.
// roleID    : The ID of a Role.
// name      : The name of the Role.
// color     : The color of the role (decimal, not hex).
// hoist     : Whether to display the role's users separately.
// perm      : The permissions for the role.
// mention   : Whether this role is mentionable
type RoleSettings struct {
	Name    string
	Color   int
	Hoist   bool
	Perm    int
	Mention bool
}

// GuildRoleCreate creates a role and edits it with the given GuildRoleSettings
// Struct requires paramater [GuildID] to be set
func (s *Session) GuildRoleCreate(guildID string, settings RoleSettings) (*discordgo.Role, error) {
	role, err := s.DG.GuildRoleCreate(guildID)
	if err != nil {
		return nil, err
	}

	// Wait until the role gets updated in the guild
	// fors.NextGuildRoleCreate().Role.ID != role.ID {
	// }

	return s.GuildRoleEdit(guildID, role.ID, settings)
}

// GuildRoleEdit edit edits the role in the given guild with 'settings'
// Struct requires parameters [GuildID] [RoleID] to be set.
func (s *Session) GuildRoleEdit(guildID, roleID string, settings RoleSettings) (*discordgo.Role, error) {
	return s.DG.GuildRoleEdit(
		guildID, roleID,
		settings.Name, settings.Color, settings.Hoist,
		settings.Perm, settings.Mention,
	)
}

// GuildRoleDelete deletes a role from a guild
func (s *Session) GuildRoleDelete(i interface{}, roleID string) error {
	guild, err := s.Guild(i)
	if err != nil {
		return err
	}

	return s.DG.GuildRoleDelete(guild.ID, roleID)
}

// GuildRoleDeleteByName deletes a role from the guild by name.
// The first argument will be used to obtain the Guild.
func (s *Session) GuildRoleDeleteByName(i interface{}, name string) error {
	guild, err := s.Guild(i)
	if err != nil {
		return err
	}

	for _, v := range guild.Roles {
		if v.Name == name {
			return s.DG.GuildRoleDelete(guild.ID, v.ID)
		}
	}

	return ErrNotFound
}

// GuildRoleDeleteByNames deletes multiple roles from the guild by name
func (s *Session) GuildRoleDeleteByNames(i interface{}, names ...string) (err error) {
	guild, err := s.Guild(i)
	if err != nil {
		return
	}

	for _, v := range names {
		err = s.GuildRoleDeleteByName(guild, v)
	}

	return
}

// SendMessage is a convenience method for sending messages to a channel
func (s *Session) SendMessage(i interface{}, text ...interface{}) (*discordgo.Message, error) {
	channelid, err := s.ChannelID(i)
	if err != nil {
		return nil, err
	}
	return s.DG.ChannelMessageSend(channelid, fmt.Sprint(text...))
}

// SendFile is a convenience method for sending files to a channel
func (s *Session) SendFile(i interface{}, filename string, rd io.Reader) (*discordgo.Message, error) {
	channelid, err := s.ChannelID(i)
	if err != nil {
		return nil, err
	}
	return s.DG.ChannelFileSend(channelid, filename, rd)
}

// SendEmbed is a convenience method for sending embeds to a channel
func (s *Session) SendEmbed(i interface{}, e interface{}) (*discordgo.Message, error) {
	var embed *discordgo.MessageEmbed
	switch t := e.(type) {
	case *discordgo.MessageEmbed:
		embed = t
	case *Embed:
		embed = t.MessageEmbed
	case string:
		embed = NewEmbed().SetDescription(t).MessageEmbed
	default:
		return nil, ErrInvalidType
	}
	channelid, err := s.ChannelID(i)
	if err != nil {
		return nil, err
	}
	return s.DG.ChannelMessageSendEmbed(channelid, embed)
}
