package irc

import (
	"github.com/justyntemme/goNotifier/web"
	"github.com/thoj/go-ircevent"
)

//SendMessage to send the ircNotification
func SendMessage(ircn web.IrcNotification) {
	ircobj := irc.IRC(ircn.From[0], ircn.From[0])
	ircobj.Connect(ircn.Server[0])
	ircobj.Join(ircn.Room[0])
	ircobj.Privmsg(ircn.To[0], ircn.Msg[0])

}
