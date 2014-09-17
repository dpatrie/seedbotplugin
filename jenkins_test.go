package seedbotplugin

import (
	"testing"
)

func TestGetJobs(t *testing.T) {
	expectedReply := "Getting jobs..."
	reply := getJobs()
	if reply != expectedReply {
		t.Errorf("Test: '%s' does not equal '%s'", reply, expectedReply)
	}
}
