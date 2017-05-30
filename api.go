package dream

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"sort"

	//Blank imports included for decoding a user's avatar into an image.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/Necroforger/discordgo"
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
func (b *Bot) MessageFromInterface(i interface{}) (*discordgo.Message, error) {
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
func (b *Bot) ChannelID(i interface{}) (string, error) {

	// Attempt to retrieve a message object from the interface
	if m, err := b.MessageFromInterface(i); err == nil {
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
func (b *Bot) GuildID(i interface{}) (string, error) {

	// Attempt to get the Message object from the interface.
	// If it fails, check the other possible types.
	if t, err := b.MessageFromInterface(i); err == nil {

		c, err := b.Channel(t.ChannelID)
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

	case *discordgo.VoiceConnection:
		return t.GuildID, nil

	case *discordgo.Member:
		return t.GuildID, nil

	default:
		return "", ErrInvalidType
	}

}

// UserID returns the userID from a variety of objects.
func (b *Bot) UserID(i interface{}) (userid string, err error) {

	if t, err := b.MessageFromInterface(i); err == nil {
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
func (b *Bot) Channel(i interface{}) (*discordgo.Channel, error) {

	// Return channel if `i` is already of type channel
	if c, ok := i.(*discordgo.Channel); ok {
		return c, nil
	}

	channelid, err := b.ChannelID(i)
	if err != nil {
		return nil, err
	}
	c, err := b.DG.State.Channel(channelid)
	if err != nil {
		c, err = b.DG.Channel(channelid)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// ChannelVoiceJoin joins the specified voice channel
func (b *Bot) ChannelVoiceJoin(guildID, channelID string, mute, deaf bool) (*discordgo.VoiceConnection, error) {
	vc, err := b.DG.ChannelVoiceJoin(guildID, channelID, mute, deaf)
	if err != nil {
		return nil, err
	}
	return vc, nil
}

// UserVoiceState finds a user's voice state from the session
func (b *Bot) UserVoiceState(userID string) (*discordgo.VoiceState, error) {

	for _, v := range b.DG.State.Guilds {
		for _, c := range v.VoiceStates {
			if c.UserID == userID {
				return c, nil
			}
		}
	}

	return nil, errors.New("VoiceState not found")
}

// UserVoiceStateJoin joins a user's voice state channel.
func (b *Bot) UserVoiceStateJoin(userID interface{}, mute, deaf bool) (*discordgo.VoiceConnection, error) {
	uid, err := b.UserID(userID)
	if err != nil {
		return nil, err
	}

	vs, err := b.UserVoiceState(uid)
	if err != nil {
		return nil, err
	}

	return b.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, mute, deaf)
}

// UserAvatarURL returns the URL of a User's Avatar. Sizes: 64, 128, 256, 512, 1024...
func (b *Bot) UserAvatarURL(userID, avatarID, size string) string {
	extension := ".jpg"
	if strings.HasPrefix(avatarID, "a_") {
		extension = ".gif"
	}
	return discordgo.EndpointCDNAvatars + userID + "/" + avatarID + extension + "?size=" + size
}

// UserAvatar returns a user's avatar decoded into an image
func (b *Bot) UserAvatar(userID, avatarID, size string) (image.Image, error) {
	url := b.UserAvatarURL(userID, avatarID, size)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// convertToOpus converts the given io.Reader stream to an Opus stream
// Using ffmpeg and dca-rs
func (b *Bot) convertToOpus(rd io.Reader) (io.Reader, error) {
	ffmpeg := exec.Command(b.Config.FfmpegPath, "-i", "pipe:0", "-f", "s16le", "-ar", "48000", "-ac", "2", "pipe:1")
	ffmpeg.Stdin = rd
	ffmpegout, err := ffmpeg.StdoutPipe()
	if err != nil {
		return nil, err
	}

	dca := exec.Command(b.Config.DcaRsPath, "--raw", "-i", "pipe:0")
	dca.Stdin = ffmpegout
	dcaout, err := dca.StdoutPipe()
	dcabuf := bufio.NewReaderSize(dcaout, 1024)
	if err != nil {
		return nil, err
	}

	err = ffmpeg.Start()
	if err != nil {
		b.Log("convertToOpus ffmpeg err: ", err)
		return nil, err
	}

	err = dca.Start()
	if err != nil {
		b.Log("convertToOpus: dca-rs err: ", err)
		return nil, err
	}

	return dcabuf, nil
}

// GuildAudioDispatcherStop a guild's currently playing audio dispatchers
func (b *Bot) GuildAudioDispatcherStop(i interface{}) error {
	guildID, err := b.GuildID(i)
	if err != nil {
		return err
	}

	disp, err := b.audioDispatcher(guildID)
	if err != nil {
		return err
	}

	disp.Stop()

	return nil
}

// GuildAudioDispatcherPause pauses the guild's currently playing audio dispatcher
func (b *Bot) GuildAudioDispatcherPause(i interface{}) error {
	guildID, err := b.GuildID(i)
	if err != nil {
		return err
	}

	disp, err := b.audioDispatcher(guildID)
	if err != nil {
		return err
	}

	disp.Pause()
	return nil
}

// GuildAudioDispatcherResume resumes the guild's currently playing audio dispatcher
func (b *Bot) GuildAudioDispatcherResume(i interface{}) error {
	guildID, err := b.GuildID(i)
	if err != nil {
		return err
	}

	disp, err := b.audioDispatcher(guildID)
	if err != nil {
		return err
	}

	disp.Resume()
	return nil
}

// GuildAudioDispatcher returns the specified guild's audio dispatcher
func (b *Bot) GuildAudioDispatcher(i interface{}) (*AudioDispatcher, error) {
	guildID, err := b.GuildID(i)
	if err != nil {
		return nil, err
	}

	return b.audioDispatcher(guildID)
}

// PlayStream plays an audio stream from the given io reader and uses ffmpeg to convert to a suitable format
func (b *Bot) PlayStream(vc *discordgo.VoiceConnection, rd io.Reader) *AudioDispatcher {
	opusStream, err := b.convertToOpus(rd)
	if err != nil {
		return nil
	}

	disp := NewAudioDispatcher(vc, opusStream)
	b.GuildAudioDispatcherStop(vc.GuildID)
	b.addAudioDispatcher(disp)

	go func() {
		disp.Start()
		// b.removeAudioDispatcher(disp.GuildID)
	}()

	return disp
}

// PlayRawStream plays a direct stream
func (b *Bot) PlayRawStream(vc *discordgo.VoiceConnection, rd io.Reader) *AudioDispatcher {
	disp := NewAudioDispatcher(vc, rd)
	b.GuildAudioDispatcherStop(vc.GuildID)
	b.addAudioDispatcher(disp)

	go func() {
		disp.Start()
		//b.removeAudioDispatcher(disp.GuildID)
	}()

	return disp
}

// PlayFile opens a file and calls PlayStream on it
func (b *Bot) PlayFile(vc *discordgo.VoiceConnection, path string) (*AudioDispatcher, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		return nil, err
	}

	disp := b.PlayStream(vc, f)
	return disp, nil
}

// PlayRawFile opens a file and calls PlayRawstream on it
func (b *Bot) PlayRawFile(vc *discordgo.VoiceConnection, path string) (*AudioDispatcher, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		return nil, err
	}

	disp := b.PlayRawStream(vc, f)
	return disp, nil
}

// Guild is a convenience method for retrieving a channel from a variety of objects
func (b *Bot) Guild(i interface{}) (*discordgo.Guild, error) {
	if g, ok := i.(*discordgo.Guild); ok {
		return g, nil
	}

	guildid, err := b.GuildID(i)
	if err != nil {
		return nil, err
	}
	guild, err := b.DG.State.Guild(guildid)
	if err != nil {
		guild, err = b.DG.Guild(guildid)
		if err != nil {
			return nil, err
		}
	}
	return guild, nil
}

// TODO Rework GuildPresence

// GuildPresence attempts to first find a guildMember object from the supplied arguments. If a member is found,
// It uses the members guildID and userID.
func (b *Bot) GuildPresence(i ...interface{}) (*discordgo.Presence, error) {
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

		gid, err := b.GuildID(i[0])
		if err != nil {
			return nil, err
		}
		uid, err := b.UserID(i[0])
		if err != nil {
			return nil, err
		}

		userID, guildID = uid, gid

	} else {

		//Get guildID from first argument
		gid, err := b.GuildID(i[0])
		if err != nil {
			return nil, err
		}

		//Get userID from second argument
		uid, err := b.UserID(i[1])
		if err != nil {
			return nil, err
		}

		userID, guildID = uid, gid
	}

	p, err := b.DG.State.Presence(guildID, userID)
	if err == nil {
		return p, nil
	}

	return nil, ErrNotFound
}

// GuildVoiceConnection returns the voice connection object for the given guild
func (b *Bot) GuildVoiceConnection(i interface{}) (*discordgo.VoiceConnection, error) {
	guildID, err := b.GuildID(i)
	if err != nil {
		return nil, err
	}

	if vc, ok := b.DG.VoiceConnections[guildID]; ok {
		return vc, nil
	}

	return nil, errors.New("Voice connection not found")
}

// GuildVoiceConnectionDisconnect finds the current guild voice connection and disconnects from it
func (b *Bot) GuildVoiceConnectionDisconnect(guildID interface{}) error {
	vc, err := b.GuildVoiceConnection(guildID)
	if err != nil {
		return err
	}

	return vc.Disconnect()
}

// GuildMember is a convenience method for fetching the member object from various sources
func (b *Bot) GuildMember(i ...interface{}) (*discordgo.Member, error) {
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
		guildid, err = b.GuildID(i[0])
		if err != nil {
			return nil, err
		}
		userid, err = b.UserID(i[0])
		if err != nil {
			return nil, err
		}
	} else {
		// If there are more than one argument present,
		// Attempt to get the guildID from the first one
		// And the userID from the second

		guildid, err = b.GuildID(i[0])
		if err != nil {
			return nil, err
		}

		userid, err = b.UserID(i[1])
		if err != nil {
			return nil, err
		}
	}

	// Fetch member from state and fall back to the restAPI if that fails.
	gm, err := b.DG.State.Member(guildid, userid)
	if err != nil {
		gm, err = b.DG.GuildMember(guildid, userid)
		if err != nil {
			return nil, err
		}
	}
	return gm, nil
}

// GuildRoles requests a guilds roles from the state or the API if not available in state
func (b *Bot) GuildRoles(i interface{}) ([]*discordgo.Role, error) {
	guild, err := b.Guild(i)
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
func (b *Bot) GuildMemberRoles(i ...interface{}) ([]*discordgo.Role, error) {
	var roles []*discordgo.Role

	member, err := b.GuildMember(i...)
	if err != nil {
		return nil, err
	}

	guildRoles, err := b.GuildRoles(member.GuildID)
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
func (b *Bot) GuildMemberRoleAdd(guildID, userID interface{}, roleID string) error {
	gid, err := b.GuildID(guildID)
	if err != nil {
		return err
	}

	uid, err := b.UserID(userID)
	if err != nil {
		return err
	}

	return b.DG.GuildMemberRoleAdd(gid, uid, roleID)
}

// GuildMemberRoleAddByName adds a role to the specified guild member
func (b *Bot) GuildMemberRoleAddByName(guildID, userID interface{}, name string) error {
	guild, err := b.Guild(guildID)
	if err != nil {
		return err
	}

	uid, err := b.UserID(userID)
	if err != nil {
		return err
	}

	for _, v := range guild.Roles {
		if v.Name == name {
			return b.DG.GuildMemberRoleAdd(guild.ID, uid, v.ID)
		}
	}

	return ErrNotFound
}

// GuildMemberRoleAddByNames ...
func (b *Bot) GuildMemberRoleAddByNames(guildID, userID interface{}, names ...string) (err error) {
	guild, err := b.Guild(guildID)
	if err != nil {
		return
	}

	uid, err := b.UserID(userID)
	if err != nil {
		return
	}

	for _, name := range names {
		err = b.GuildMemberRoleAddByName(guild, uid, name)
	}
	return
}

// GuildMemberRoleRemove removes the specified role to a given member
//  guildID   : The ID of a Guild.
//  userID    : The ID of a User.
//  roleID 	  : The ID of a Role to be removed from the user.
func (b *Bot) GuildMemberRoleRemove(guildID, userID interface{}, roleID string) error {
	gid, err := b.GuildID(guildID)
	if err != nil {
		return err
	}

	uid, err := b.UserID(userID)
	if err != nil {
		return err
	}

	return b.DG.GuildMemberRoleRemove(gid, uid, roleID)
}

// GuildMemberRoleRemoveByName removes a role from a member by name
func (b *Bot) GuildMemberRoleRemoveByName(guildID, userID interface{}, rolename string) error {
	gid, err := b.GuildID(guildID)
	if err != nil {
		return err
	}

	uid, err := b.UserID(userID)
	if err != nil {
		return err
	}

	memberRoles, err := b.GuildMemberRoles(gid, uid)
	if err != nil {
		return err
	}

	for _, v := range memberRoles {
		if v.Name == rolename {
			return b.GuildMemberRoleRemove(gid, uid, v.ID)
		}
	}

	return ErrNotFound
}

// GuildMemberRolesRemoveByName removes a list of roles by name from a guild member
func (b *Bot) GuildMemberRolesRemoveByName(guildID, userID interface{}, rolenames ...string) (err error) {
	gid, err := b.GuildID(guildID)
	if err != nil {
		return err
	}

	uid, err := b.UserID(userID)
	if err != nil {
		return err
	}

	for _, rolename := range rolenames {
		err = b.GuildMemberRoleRemoveByName(gid, uid, rolename)
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
	GuildID string
	RoleID  string
	Name    string
	Color   int
	Hoist   bool
	Perm    int
	Mention bool
}

// GuildRoleCreate creates a role and edits it with the given GuildRoleSettings
// Struct requires paramater [GuildID] to be set
func (b *Bot) GuildRoleCreate(settings RoleSettings) (*discordgo.Role, error) {
	role, err := b.DG.GuildRoleCreate(settings.GuildID)
	if err != nil {
		return nil, err
	}

	// Wait until the role gets updated in the guild
	// for b.NextGuildRoleCreate().Role.ID != role.ID {
	// }

	settings.RoleID = role.ID
	return b.GuildRoleEdit(settings)
}

// GuildRoleEdit edit edits the role in the given guild with 'settings'
// Struct requires parameters [GuildID] [RoleID] to be set.
func (b *Bot) GuildRoleEdit(settings RoleSettings) (*discordgo.Role, error) {
	return b.DG.GuildRoleEdit(
		settings.GuildID, settings.RoleID,
		settings.Name, settings.Color, settings.Hoist,
		settings.Perm, settings.Mention,
	)
}

// GuildRoleDelete deletes a role from a guild
func (b *Bot) GuildRoleDelete(i interface{}, roleID string) error {
	guild, err := b.Guild(i)
	if err != nil {
		return err
	}

	return b.DG.GuildRoleDelete(guild.ID, roleID)
}

// GuildRoleDeleteByName deletes a role from the guild by name.
// The first argument will be used to obtain the Guild.
func (b *Bot) GuildRoleDeleteByName(i interface{}, name string) error {
	guild, err := b.Guild(i)
	if err != nil {
		return err
	}

	for _, v := range guild.Roles {
		if v.Name == name {
			return b.DG.GuildRoleDelete(guild.ID, v.ID)
		}
	}

	return ErrNotFound
}

// GuildRoleDeleteByNames deletes multiple roles from the guild by name
func (b *Bot) GuildRoleDeleteByNames(i interface{}, names ...string) (err error) {
	guild, err := b.Guild(i)
	if err != nil {
		return
	}

	for _, v := range names {
		err = b.GuildRoleDeleteByName(guild, v)
	}

	return
}

// SendMessage is a convenience method for sending messages to a channel
func (b *Bot) SendMessage(i interface{}, text ...interface{}) (*discordgo.Message, error) {
	channelid, err := b.ChannelID(i)
	if err != nil {
		return nil, err
	}
	return b.DG.ChannelMessageSend(channelid, fmt.Sprint(text...))
}

// SendFile is a convenience method for sending files to a channel
func (b *Bot) SendFile(i interface{}, filename string, rd io.Reader) (*discordgo.Message, error) {
	channelid, err := b.ChannelID(i)
	if err != nil {
		return nil, err
	}
	return b.DG.ChannelFileSend(channelid, filename, rd)
}

// SendEmbed is a convenience method for sending embeds to a channel
func (b *Bot) SendEmbed(i interface{}, e interface{}) (*discordgo.Message, error) {
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
	channelid, err := b.ChannelID(i)
	if err != nil {
		return nil, err
	}
	return b.DG.ChannelMessageSendEmbed(channelid, embed)
}
