package dream

import (
	"io"
	"os"
	"sync"

	"github.com/Necroforger/discordgo"
)

//Generate the AddHandler and AddHandlerOnce methods
//go:generate go run tools/addhandlers/main.go

//Generate nextEvent functions
//go:generate go run tools/nextevent/main.go

//Session contains session information.
type Session struct {
	sync.Mutex

	// LogOutput is the Writer where all the log events are sent over.
	// It is set to os.Stdout by default.
	LogOutput io.Writer

	DG     *discordgo.Session
	Config Config

	// Client is the bot's user
	Client *discordgo.User

	// AudioDispatchers stores the Audio Dispatchers belonging to each guild
	AudioDispatchers map[string]*AudioDispatcher
}

//Config represents dream's configuration
type Config struct {

	// FfmpegPath is the path to use for the ffmpeg command.
	// Default "ffmpeg"
	FfmpegPath string

	//DcaRsPath is the location of the dca-rs executeable for encoding opus
	// Default: "./dca-rs"
	DcaRsPath string
}

//NewConfig returns the default configuration options for the bot
func NewConfig() Config {
	return Config{
		FfmpegPath: "ffmpeg",
		DcaRsPath:  "./dca-rs",
	}
}

//New returns a new Session.
func New(conf Config, args ...interface{}) (*Session, error) {
	bot := &Session{}
	bot.Config = conf
	bot.AudioDispatchers = map[string]*AudioDispatcher{}
	bot.LogOutput = os.Stdout

	session, err := discordgo.New(args...)
	if err != nil {
		return nil, err
	}
	bot.DG = session

	return bot, nil
}

func (s *Session) addAudioDispatcher(ad *AudioDispatcher) {
	s.Lock()
	s.AudioDispatchers[ad.GuildID] = ad
	s.Unlock()
}

// removeAudioDispatcher removes the audio dispatcher from the map with ID guildID
func (s *Session) removeAudioDispatcher(guildID string) {
	s.Lock()
	delete(s.AudioDispatchers, guildID)
	s.Unlock()
}

// audioDispatcher returns an audio dispatcher by guild ID
func (s *Session) audioDispatcher(guildID string) (*AudioDispatcher, error) {
	s.Lock()
	defer s.Unlock()

	if v, ok := s.AudioDispatchers[guildID]; ok {
		return v, nil
	}
	return nil, ErrNotFound
}

// Open connects to discord websockets
func (s *Session) Open() error {

	err := s.DG.Open()
	if err != nil {
		return err
	}

	return nil
}
