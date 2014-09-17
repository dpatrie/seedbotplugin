package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
	"strings"
)

type Echo struct{}

func (p Echo) Name() string {
	return "Echo v1.0"
}

func (p Echo) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {
	if strings.HasPrefix(msg.Body(), "echo ") {
		if msg.From() != bot.FullName() {
			reply := echo(msg.Body())
			bot.Send(reply)
		}
	}
	return nil
}

func echo(message string) string {
	ns := strings.Replace(message, "echo ", "", 1)
	return ns
}
