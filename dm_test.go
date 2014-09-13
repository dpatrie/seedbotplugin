package seedbotplugin

import (
	"testing"
)

var dmTests = []struct {
	in  string
	out string
}{
	{
		in:  "bot",
		out: "my ears are burning",
	},
	{
		in:  "bot plus some stuff",
		out: "my ears are burning",
	},
	{
		in:  "some stuff and then bot later",
		out: "my ears are burning",
	},
	{
		in:  "some stuff and then bot",
		out: "my ears are burning",
	},
	{
		in:  "some stuff and then that's it",
		out: "",
	},
}

func TestDm(t *testing.T) {
	for i, test := range dmTests {
		actual := dm(test.in, "bot")
		if actual != test.out {
			t.Errorf("Test %d: '%s' does not equal '%s'", i, actual, test.out)
		}
	}
}
