package radical

import (
	"testing"
)

func TestToFrom(t *testing.T) {
	list := [][2]int32{
		{1, 2},
		{-3, 4},
		{5, -6},
		{-7, -8},
		{39, -110},
	}
	for _, item := range list {
		a, b := item[0], item[1]
		t.Logf("check %d and %d\n", a, b)
		c := To(a, b)
		a1, b1 := From(c)
		t.Logf("is now %d and %d\n", a1, b1)
		if a != a1 {
			t.Errorf("A want: %d -- got:%d", a, a1)
		}
		if b != b1 {
			t.Errorf("B want: %d -- got:%d", a, a1)
		}
	}
}

type ranger struct {
	lat1, lon1 int32
	lat2, lon2 int32
	lat3, lon3 int32
	within     bool
}

func TestBetween(t *testing.T) {
	list := []ranger{
		{10, 100, 20, 150, 14, 125, true},
		{10, 100, 20, 150, 24, 125, false},
		{20, -100, 22, -80, 21, -90, true},
	}
	for _, item := range list {
		a := To(item.lat1, item.lon1)
		b := To(item.lat2, item.lon2)
		c := To(item.lat3, item.lon3)
		//t.Logf("check %d and %d contains %d (%t)\n", item.from, item.to, item.check, item.within)
		betwixt := Between(a, b, c)
		if betwixt != item.within {
			t.Errorf("failed check {%d,%d} and {%d,%d} contains {%d,%d} (%t)\n",
				item.lat1, item.lon1,
				item.lat2, item.lon2,
				item.lat3, item.lon3,
				item.within,
				/*
				 */
			)
		}
	}
}
