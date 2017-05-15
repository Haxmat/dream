package dream

import "github.com/Necroforger/discordgo"

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
