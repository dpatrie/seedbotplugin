package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
	"time"
)

var rickroll = []string{
	"I just wanna tell you how I'm feeling",
	"Gotta make you understand",
	"Never gonna give you up",
	"Never gonna let you down",
	"Never gonna run around and desert you",
	"Never gonna make you cry",
	"Never gonna say goodbye",
	"Never gonna tell a lie and hurt you",
}

type RickRoll struct{}

func (p RickRoll) Name() string {
	return "RickRoll v1.0"
}

func (p RickRoll) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {
	if msg.Body() == bot.Name()+": how are you feeling?" {
		time.Sleep(5 * time.Second)
		for _, s := range rickroll {
			bot.Send(s)
			time.Sleep(3 * time.Second)
		}
	}
	return nil
}
