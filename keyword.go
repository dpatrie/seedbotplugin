// Copyright 2014 Gabriel Guzman <gabe@lifewaza.com>
// All rights reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

// Package keyword gives the bot the ability to analyze messages
//  to determine word counts
package seedbotplugin

import (
	"bitbucket.org/gabeguz/gobothelper"
	"fmt"
	"github.com/seedboxtech/kwextractor"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// Helper struct that will implement the helper interface
type Helper struct {
}

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
func (h *Helper) Send(message gobothelper.Message, cb gobothelper.Bot) {
	if strings.Contains(message.Body, "keywords") {
		kws := keywords("/tmp/chatlog")
		if len(kws) > 0 {
			for i := 0; i < len(kws) && i < 10; i++ {
				text := fmt.Sprintf("k: %s v: %d", kws[i].Word, kws[i].Count)
				cb.Reply(text)
			}
		} else {
			log.Printf("No keywords available.")
		}
	}
}

func keywords(logPath string) keywordSlice {
	fileContents, err := ioutil.ReadFile(logPath)
	if err != nil {
		log.Fatal(err)
	}
	kws := kwextractor.KeywordsAndFrequencyFrom(string(fileContents))
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
