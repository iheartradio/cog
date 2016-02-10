package clog

import (
	"encoding/json"

	"github.com/thatguystone/cog/config"
)

// JSONFormat formats messages as JSON
type JSONFormat struct{}

func init() {
	RegisterFormatter("JSON",
		func(args config.Args) (Formatter, error) {
			return JSONFormat{}, nil
		})
}

// FormatEntry implements Formatter
func (JSONFormat) FormatEntry(e Entry) ([]byte, error) {
	return json.Marshal(e)
}
