package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
	"log"
	"testing"
)

type Bot struct{}

func (p Bot) Connect() error {
	return nil
}

func (p Bot) Listen() chan xmppbot.Message {
	return nil
}

func (p Bot) Log(string) {
}

func (p Bot) Name() string {
	return "bot"
}

func (p Bot) FullName() string {
	return "testbot"
}

func (p Bot) Send(string) {
}

func (p Bot) SetLogger(*log.Logger) {
}

type Message struct{}

func (m Message) Body() string {
	return "this is a test of the emergency broadcast network."
}

func (m Message) From() string {
	return "nottestbot"
}

func TestEcho(t *testing.T) {
	message := Message{}
	bot := Bot{}
	p := Echo{}
	p.Execute(message, bot)
	expectedReply := "this is a test of the emergency broadcast network."
	reply := echo("echo this is a test of the emergency broadcast network.")
	if reply != expectedReply {
		t.Errorf("Test: '%s' does not equal '%s'", reply, expectedReply)
	}
}
