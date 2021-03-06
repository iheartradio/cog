package check

import "testing"

func TestErrorerBasic(t *testing.T) {
	c := New(t)

	er := Errorer{}
	c.Error(er.Err())
	c.True(er.Fail())
}

func TestErrorerIgnoreTestFns(t *testing.T) {
	c := New(t)

	er := Errorer{
		IgnoreTestFns: true,
	}

	c.NotError(er.Err())
	c.False(er.Fail())
}

func testErrorerOnlyInHere(c *C, er *Errorer) {
	c.Error(er.Err())
}

func testErrorerOnlyInNotHere(c *C, er *Errorer) {
	c.NotError(er.Err())
}

func TestErrorerOnlyIn(t *testing.T) {
	c := New(t)

	er := Errorer{
		OnlyIn: []string{"testErrorerOnlyInHere"},
	}

	testErrorerOnlyInHere(c, &er)
	testErrorerOnlyInNotHere(c, &er)
}

func testErrorerSameCodePath(c *C, er *Errorer, fail bool) {
	c.Equal(fail, er.Fail())
}

func TestErrorerSameCodePath(t *testing.T) {
	c := New(t)

	er := Errorer{}
	for i := 0; i < 5; i++ {
		testErrorerSameCodePath(c, &er, i == 0)
	}
}
