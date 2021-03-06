// Copyright 2014 Gabriel Guzman <gabe@lifewaza.com>
// All rights reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

// Package keyword gives the bot the ability to analyze messages
//  to determine word counts
package seedbotplugin

import (
	"fmt"
    "github.com/seedboxtech/xmppbot"
	"github.com/seedboxtech/kwextractor"
	"log"
	"sort"
	"strings"
)

// Helper struct that will implement the helper interface
type Keyword struct {
}

func (p Keyword) Name() string {
	return "Keyword v1.0"
}

type messages []string

type keywordSlice []keyword

type keyword struct {
	Count int
	Word  string
}

// Implement the sort interface
func (k keywordSlice) Len() int {
	return len(k)
}

func (k keywordSlice) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k keywordSlice) Less(i, j int) bool {
	return k[i].Count > k[j].Count
}

// Send allows the bot to send a message to this helper
func (p Keyword) Execute(message xmppbot.Message, cb xmppbot.Bot) error {

    messageBuffer := make(messages, 0, 0)
    messageBuffer = append(messageBuffer, message.Body())
	if strings.Contains(message.Body(), "keywords") {
		kws := keywords(messageBuffer)
		if len(kws) > 0 {
			for i := 0; i < len(kws) && i < 10; i++ {
				text := fmt.Sprintf("k: %s v: %d", kws[i].Word, kws[i].Count)
				cb.Send(text)
			}
		} else {
			log.Printf("No keywords available.")
		}
	}
	return nil
}

func keywords(mb []string) keywordSlice {
	kws := kwextractor.KeywordsAndFrequencyFrom(mb)
	s := make(keywordSlice, 0, len(kws))

	for k, v := range kws {
		t := &keyword{
			Count: v,
			Word:  k,
		}
		s = append(s, *t)
	}
	sort.Sort(s)
	return s
}
