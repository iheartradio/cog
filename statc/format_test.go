package statc

import (
	"fmt"
	"testing"

	"github.com/thatguystone/cog/check"
	"github.com/thatguystone/cog/cio/eio"
)

type formatErrors struct{}

func init() {
	RegisterFormatter("errors",
		func(eio.Args) (Formatter, error) {
			return formatErrors{}, nil
		})
}

func (formatErrors) FormatSnap(Snapshot) ([]byte, error) {
	return nil, fmt.Errorf("i have issues with that snapshot")
}

func (formatErrors) MimeType() string {
	return "application/errors"
}

func TestFormatErrors(t *testing.T) {
	c := check.New(t)

	c.Panic(func() {
		RegisterFormatter("json", nil)
	})

	_, err := newFormatter("iDontExist", nil)
	c.Error(err)
}
