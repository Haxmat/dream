package dream

import (
	"fmt"

	"github.com/Necroforger/discordgo"
)

const (
	eMsg         = "MSG"
	eMsgCmd      = "CMD"
	eMsgDel      = "MSGDEL"
	eMsgUpdate   = "MSGUPDATE"
	eMsgReact    = "MSGREACT"
	eMsgReactDel = "MSGREACTDEL"
)

func (b *Bot) onReady(s *discordgo.Session, r *discordgo.Ready) {
	if r.User == nil {
		return
	}
	b.Log(0, fmt.Sprintf("Dream Bot successfully connected as user [%s] and is now serving in [%d] guilds", r.User.Username, len(r.Guilds)))
}
