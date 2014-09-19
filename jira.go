package seedbotplugin

import (
	"fmt"
	"github.com/seedboxtech/libgojira"
	"github.com/seedboxtech/xmppbot"
	"strings"
)

const (
	SPRINT_ITEMS = "sprint"
)

type Jira struct {
	client *libgojira.JiraClient
}

func (p Jira) Name() string {
	return "Jira v1.0"
}

func (p Jira) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {

	if p.client == nil {
		//TODO: Read options from file??

		p.client = libgojira.NewJiraClient(libgojira.Options{
			User:       "YOUWISH",
			Passwd:     "NONONO",
			NoCheckSSL: true,
			Server:     "jira.gammae.com",
			Projects:   []string{"ADX", "MTA"},
		})
	}

	switch parseMsg(msg.Body()) {
	case SPRINT_ITEMS:
		res, err := p.client.Search(&libgojira.SearchOptions{
			Projects:      []string{"ADX", "MTA"},
			CurrentSprint: true,
			NotType:       []string{"Sub-task"},
		})
		if err != nil {
			return err
		}
		bot.Send(render(res))
	}

	return nil
}

func render(issues []*libgojira.Issue) string {
	out := "\n"

	for _, i := range issues {
		out += fmt.Sprintf("%s \t %s \t %s \n", i.Key, i.Status, i.Summary)
	}

	return out
}

func parseMsg(msg string) string {
	if strings.HasPrefix(msg, "jira ") {
		return msg[5:]
	}
	return ""
}
