package seedbotplugin

import (
	"code.google.com/p/go.net/html"
	"github.com/seedboxtech/xmppbot"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Url struct{}

func (p Url) Name() string {
	return "Url v1.0"
}

func (p Url) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {

	u, err := url.Parse(msg.Body())
	if err != nil {
		return err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return (err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return (err)
	}
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return (err)
	}

	var g func(*html.Node)
	g = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			for _, value := range n.Attr {
				bot.Send(value.Val)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			g(c)
		}
	}
	g(doc)
	return nil
}
