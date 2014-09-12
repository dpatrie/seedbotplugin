package seedbotplugin

import (
	"testing"
)

func TestCodeQuote(t *testing.T) {
	result := codeQuote()
	if len(result) == 0 {
		t.Errorf("No valid quote returned.")
	}
}

func TestIntegrationQuote(t *testing.T) {
	result := integrationQuote()
	if len(result) == 0 {
		t.Errorf("No valid integration quote returned.")
	}
}

func TestAdminQuote(t *testing.T) {
	result := integrationQuote()
	if len(result) == 0 {
		t.Errorf("No valid integration quote returned.")
	}
}
