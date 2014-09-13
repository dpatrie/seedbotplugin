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
            ns := strings.Replace(msg.Body(), "echo ", "", 1)
		    bot.Send(ns)
	    }
    }
	return nil
}
