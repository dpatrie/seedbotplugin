package seedbotplugin

import (
	"testing"
)

func TestEcho(t *testing.T) {
	expectedReply := "this is a test of the emergency broadcast network."
	reply := echo("echo this is a test of the emergency broadcast network.")
	if reply != expectedReply {
		t.Errorf("Test: '%s' does not equal '%s'", reply, expectedReply)
	}
}
