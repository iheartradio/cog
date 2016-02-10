package clog

import (
	"testing"
	"time"

	"github.com/thatguystone/cog/check"
	"github.com/thatguystone/cog/config"
)

func TestHumanFormatNew(t *testing.T) {
	c := check.New(t)

	_, err := newFormatter(FormatterConfig{
		Name: "Human",
		Args: config.Args{
			"ShortTime": "funny",
		},
	})
	c.MustError(err)
}

func TestHumanFormatCoverage(t *testing.T) {
	check.New(t)

	f := HumanFormat{}
	f.FormatEntry(Entry{
		Time: time.Now(),
		Data: Data{
			"test": 1,
		},
	})

	f.Args.ShortTime = true
	f.FormatEntry(Entry{
		Time: time.Now(),
	})
}
