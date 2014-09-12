package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
)

type Plugin interface {
	Name() string
	Execute(msg xmppbot.Message, bot xmppbot.Bot) error
}
