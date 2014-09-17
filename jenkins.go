package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
	"strings"
)

type Jenkins struct{}

func (p Jenkins) Name() string {
	return "Jenkins v1.0"
}

func (p Jenkins) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {
	var status string
	if strings.HasPrefix(msg.Body(), "jenkins ") {
		if msg.From() != bot.FullName() {
			if strings.Contains("jenkins build", msg.Body()) {
				status = build()
			}
			bot.Send(status)
		}
	}
	return nil
}

func getJobs() string {
	return "Getting jobs..."
}

func build() string {
	return "Build started..."
}
