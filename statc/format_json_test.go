package statc

import (
	"testing"

	"github.com/thatguystone/cog/check"
	"github.com/thatguystone/cog/cio/eio"
)

func TestJSONFormatNew(t *testing.T) {
	c := check.New(t)

	jf, err := newFormatter("json", eio.Args{
		"pretty": true,
	})
	c.NotError(err)
	c.True(jf.(JSONFormat).Args.Pretty)
}

func TestJSONFormatEmpty(t *testing.T) {
	c := check.New(t)

	b, err := JSONFormat{}.Format(Snapshot{})
	c.MustNotError(err)

	c.Equal(string(b), `{}`)
}

func TestJSONFormatBasic(t *testing.T) {
	c := check.New(t)

	snap := Snapshot{}
	snap.addTestData()
	snap.Add(newName("a.really.nested.value"), int64(1))

	b, err := JSONFormat{}.Format(snap)
	c.MustNotError(err)

	c.Equal(
		string(b),
		`{"a":{"really":{"nested":{"value":1}}},"bool":{"false":false,"true":true},"float":1.2445,"int":1,"str":"string"}`)
}

func TestJSONFormatPretty(t *testing.T) {
	c := check.New(t)

	snap := Snapshot{}
	snap.addTestData()
	snap.Add(newName("a.really.nested.value"), int64(1))

	jf := JSONFormat{}
	jf.Args.Pretty = true

	b, err := jf.Format(snap)
	c.MustNotError(err)

	c.Log(string(b))

	out := `{
	"a": {
		"really": {
			"nested": {
				"value": 1
			}
		}
	},
	"bool": {
		"false": false,
		"true": true
	},
	"float": 1.2445,
	"int": 1,
	"str": "string"` + "\n}"

	c.Equal(string(b), out)
}

func TestJSONFormatErrors(t *testing.T) {
	c := check.New(t)

	snap := Snapshot{}
	snap.Add(newName("test"), struct{}{})

	jf := JSONFormat{}
	jf.Args.Pretty = true

	_, err := jf.Format(snap)
	c.Error(err)
}

func BenchmarkJSONFormat(b *testing.B) {
	c := check.New(b)

	snap := Snapshot{}
	snap.addTestData()

	for i := 0; i < b.N; i++ {
		_, err := JSONFormat{}.Format(snap)
		c.MustNotError(err)
	}
}
