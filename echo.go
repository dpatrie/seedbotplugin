package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
)

type Echo struct{}

func (p Echo) Name() string {
	return "Echo v1.0"
}
func (p Echo) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {
	if msg.From() != bot.FullName() {
		bot.Send(msg.Body())
	}
	return nil
}
