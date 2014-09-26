package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
    "net/url"
    "log"
)

type Url struct{}

func (p Url) Name() string {
	return "Url v1.0"
}

func (p Url) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {

    url, err := url.Parse(msg.Body())
    if err != nil {
        return err
    }

    log.Println(url)

	return nil
}
