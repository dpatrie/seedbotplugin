// Copyright 2014 Gabriel Guzman <gabe@lifewaza.com>
// All rights reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

// Package chatlog logs all chat messages to a text file
package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
	"log"
	"os"
)

// Helper struct that will implement the Helper interface
type ChatLog struct {
}

func (c ChatLog) Name() string {
	return "Chatlog v1.0"
}

// Send allows the bot to send a message to this helper
func (c ChatLog) Execute(message xmppbot.Message, bot xmppbot.Bot) error {
	f := getLogFileHandle("/tmp/chatlog")
	logit(message.Body(), f)
	return nil
}

func getLogFileHandle(path string) (f *os.File) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return f
}

func logit(message string, f *os.File) {
	log.SetOutput(f)
	log.Println(message)
	defer f.Close()
}
