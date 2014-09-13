// Copyright 2014 Gabriel Guzman <gabe@lifewaza.com>
// All rights reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

// Package stathat gives the bot the ability to send info to stathat.com
package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
	"github.com/stathat/go"
)

// Helper struct that will implement the helper interface
type StatHat struct {
}

func (p StatHat) Name() string {
	return "StatHat v1.0"
}

// Send allows the bot to send a message to this helper
func (p StatHat) Execute(message xmppbot.Message, bot xmppbot.Bot) error {
	stathat.PostEZCount("chat seen", "gguzman.work@gmail.com", 1)
	return nil
}
