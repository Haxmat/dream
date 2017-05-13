package dream

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/Necroforger/discordgo"
)

//Generate the AddHandler and AddHandlerOnce methods
//go:generate go run tools/addhandlers/main.go

//Generate the listener functions for logging.
//go:generate go run tools/listeners/main.go

//Generate nextEvent functions
//go:generate go run tools/nextevent/main.go

//Bot contains all information relating to dream bot.
type Bot struct {
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

	// LogConfig is the configuration for what events dream should be logging.
	// Dream encodes any events requested from the LogConfig and outputs them to Bot.LogOutput
	LogConfig LoggingConfig
}

//NewConfig returns the default configuration options for the bot
func NewConfig() Config {
	return Config{
		LogConfig:  LoggingConfig{},
		FfmpegPath: "ffmpeg",
		DcaRsPath:  "./dca-rs",
	}
}

//New returns a new Bot object.
func New(conf Config, args ...interface{}) (*Bot, error) {
	bot := &Bot{}
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

// Log prints various information to the console based on the current LogLevel
func (b *Bot) Log(data ...interface{}) {
	fmt.Fprintln(b.LogOutput, data...)
}

func (b *Bot) addAudioDispatcher(ad *AudioDispatcher) {
	b.Lock()
	b.AudioDispatchers[ad.GuildID] = ad
	b.Unlock()
}

// removeAudioDispatcher removes the audio dispatcher from the map with ID guildID
func (b *Bot) removeAudioDispatcher(guildID string) {
	b.Lock()
	delete(b.AudioDispatchers, guildID)
	b.Unlock()
}

// audioDispatcher returns an audio dispatcher by guild ID
func (b *Bot) audioDispatcher(guildID string) (*AudioDispatcher, error) {
	b.Lock()
	defer b.Unlock()

	if v, ok := b.AudioDispatchers[guildID]; ok {
		return v, nil
	}
	return nil, ErrNotFound
}

// Adds the bot's default message handlers
func (b *Bot) addHandlers() {
	//Add the bot's build in handlers
	b.DG.AddHandler(b.onReady)

	//Add the listeners that the LogConfig requests
	b.registerListeners(b.Config.LogConfig)
}

// Open adds the bot's default handlers and begins listening for messages
func (b *Bot) Open() error {
	b.addHandlers()

	// Set the bot's default user
	user, err := b.DG.User("@me")
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user object nil")
	}
	b.Client = user

	//Connect to discord
	err = b.DG.Open()
	if err != nil {
		b.Log(0, "Error opening dream session: "+fmt.Sprint(err))
		return err
	}

	return nil
}
