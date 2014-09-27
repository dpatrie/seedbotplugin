package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
	"log"
	"net/url"
)

type Url struct{}

func (p Url) Name() string {
	return "Url v1.0"
}

func (p Url) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {

	u, err := url.Parse(msg.Body())
	if err != nil {
		return err
	}

	log.Printf("host: %s\n", u.Host)

	return nil
}
