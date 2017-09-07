package dream

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

/*

Util provides various utility functions

*/

// StatusColor returns the colour associated with an online status
func StatusColor(status discordgo.Status) int {
	color := 0
	switch status {
	case discordgo.StatusDoNotDisturb:
		color = 0xff0000
	case discordgo.StatusOnline:
		color = 0x00ff00
	case discordgo.StatusIdle:
		color = 0xffff00
	}
	return color
}

// CreationTime returns the creation time of a Snowflake ID relative to the creation of Discord.
// Taken from https://github.com/Moonlington/FloSelfbot/blob/master/commands/commandutils.go#L117
func CreationTime(ID string) (t time.Time, err error) {
	i, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return
	}
	timestamp := (i >> 22) + 1420070400000
	t = time.Unix(timestamp/1000, 0)
	return
}
