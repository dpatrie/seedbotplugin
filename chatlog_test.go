package seedbotplugin

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestLogit(t *testing.T) {
	testString := "This is a string to log."
	f, err := ioutil.TempFile("", "chatlogtest")
	if err != nil {
		t.Errorf("error opening file: %v", err)
	}
	defer f.Close()
	logit(testString, f)

	contents, err := ioutil.ReadFile(f.Name())
	if err != nil {
		t.Errorf("error reading file")
	}
	if !strings.Contains(string(contents), testString) {
		t.Errorf("'%s' does not equal '%s'", testString, string(contents))
	}

	os.Remove(f.Name())
}
