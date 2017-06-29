package dream

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/Necroforger/discordgo"
)

// TODO: Add error handling

const (
	audioPlay = iota
	audioPause
	audioStop
	audioResume
)

// Error values
var (
	ErrVoiceConnectionNil = errors.New("err: voice connection is nil")
	ErrAlreadyPlaying     = errors.New("err: already playing")
	ErrControlChannel     = errors.New("err: Reading from control channel failed")
)

var (
	//ErrTimedOut is returned when an opus packet took too long to send
	ErrTimedOut = errors.New("timed out")
)

//AudioDispatcher AudioDispatcher accepts an opus stream and a voice connection. Calling start will play over the voice connection.
type AudioDispatcher struct {
	sync.Mutex
	VC      *discordgo.VoiceConnection
	playing bool //True when not paused
	stopped bool //True when the stop control is sent

	// Control is used to stop/play/resume the currently playing audio
	control chan int

	// Source is a stream of opus data
	source io.Reader

	GuildID   string
	ChannelID string

	// Duration is the duration the AudioDispatcher has been playing for.
	// It will account for time paused.
	Duration time.Duration
}

// NewAudioDispatcher Creates a new audio dispatcher given a  `VoiceConnection`[vc] and an `io.Reader`[source]
// The io.Reader must be a stream of opus data.
func NewAudioDispatcher(vc *discordgo.VoiceConnection, source io.Reader) *AudioDispatcher {
	return &AudioDispatcher{
		VC:        vc,
		GuildID:   vc.GuildID,
		ChannelID: vc.ChannelID,
		playing:   false,
		control:   make(chan int),
		source:    source,
	}
}

//IsStopped returns if the player is stopped or not
func (a *AudioDispatcher) IsStopped() bool {
	return a.stopped
}

//Resume Resumes the currently playing audio
func (a *AudioDispatcher) Resume() {
	a.Lock()
	if !a.playing && !a.stopped {
		a.control <- audioResume
		a.playing = true
	}
	a.Unlock()
}

//Pause Pauses the currently playing audio
func (a *AudioDispatcher) Pause() {
	a.Lock()
	if a.playing && !a.stopped {
		a.control <- audioPause
		a.playing = false
	}
	a.Unlock()
}

//Stop Stops the currently playing audio and ends the dispatcher
func (a *AudioDispatcher) Stop() {
	a.Lock()
	if !a.stopped {
		a.control <- audioStop
		a.stopped = true
	}
	a.Unlock()
}

// Wait Waits for the player to finish
func (a *AudioDispatcher) Wait() {
	for !a.stopped {
		time.Sleep(time.Millisecond * 500)
	}
}

//Start starts playing audio on the given voice channel
func (a *AudioDispatcher) Start() (err error) {
	if a.VC == nil {
		return ErrVoiceConnectionNil
	}
	if a.playing {
		return ErrAlreadyPlaying
	}

	a.Lock()
	a.playing = true
	a.Unlock()

	a.VC.Speaking(true)
	defer a.VC.Speaking(false)

	defer func() {
		a.Lock()
		a.stopped = true
		a.Unlock()
	}()

	startTime := time.Now()

	updateDuration := func() {
		// a.Lock()
		a.Duration = time.Now().Sub(startTime)
		// a.Unlock()
	}

	for {
		select {
		case cmd := <-a.control:
			switch cmd {
			case audioPause:
				pausedStart := time.Now()
				for {
					v, ok := <-a.control
					if !ok {
						return
					}
					if v == audioResume {
						startTime = startTime.Add(time.Now().Sub(pausedStart))
						break
					}
					if v == audioStop {
						startTime = startTime.Add(time.Now().Sub(pausedStart))
						updateDuration()
						return nil
					}
				}
			case audioStop:
				updateDuration()
				return nil
			}
		default:
		}

		if a.VC == nil {
			fmt.Println("AudioDispatcher: ERR Voice connection became nil")
			break
		}

		opus, err := readOpus(a.source)
		if err != nil {
			if err == io.ErrUnexpectedEOF || err == io.EOF {
				return nil // This is normal, it is the end of the file
			}
			fmt.Println("AudioDispatcher: ", err)
		}

		select {
		case a.VC.OpusSend <- opus:
		case <-time.After(time.Second * 1):
			fmt.Println("AudioDispatcher: OpusSend timed out")
			return ErrTimedOut
		}

		updateDuration()
	}
	return nil
}

func readOpus(source io.Reader) ([]byte, error) {
	var opuslen int16
	err := binary.Read(source, binary.LittleEndian, &opuslen)
	if err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil, err
		}
		return nil, errors.New("ERR reading opus header")
	}

	var opusframe = make([]byte, opuslen)
	err = binary.Read(source, binary.LittleEndian, &opusframe)
	if err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil, err
		}
		return nil, errors.New("ERR reading opus frame")
	}

	return opusframe, nil
}
