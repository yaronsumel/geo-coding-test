package list

import (
	"reflect"
	"testing"
	"github.com/yaronsumel/geo-coding-test/src/place"
)

// TestAppend
func TestAppend(t *testing.T) {
	x := PlacesList{}
	p := &place.Place{Id: 1}
	x.Append(p)
	if !reflect.DeepEqual(x[0], p) {
		t.FailNow()
	}
	x2 := PlacesList{}
	x2.Append(p)
	x2.Append(&place.Place{Distance: 0.2134})
	x2.Append(&place.Place{Distance: 2341234.2134})
	for k := 0; k < 11; k++ {
		p := &place.Place{Id: k, Distance: float64(k)}
		x2.Append(p)
	}
}

// TestDeleteI
func TestDeleteI(t *testing.T) {
	x := PlacesList{}
	x = append(x, &place.Place{Id: 1}, &place.Place{Id: 2})
	x.DeleteI(0)
	if x[0].Id == 1 {
		t.FailNow()
	}
}

// TestTopFive
func TestTopFive(t *testing.T) {
	x := PlacesList{}
	x = append(x, &place.Place{Id: 1}, &place.Place{Id: 2}, &place.Place{Id: 3}, &place.Place{Id: 4}, &place.Place{Id: 5})
	if len(x.TopFive()) != 5 {
		t.FailNow()
	}
	x2 := PlacesList{}
	x2 = append(x2, &place.Place{Id: 4}, &place.Place{Id: 5})
	if len(x2.TopFive()) != 2 {
		t.FailNow()
	}
}

// TestTopFive
func TestBottomFive(t *testing.T) {
	x := PlacesList{}
	x = append(x, &place.Place{Id: 1}, &place.Place{Id: 2}, &place.Place{Id: 3}, &place.Place{Id: 4}, &place.Place{Id: 5})
	b5 := x.BottomFive()
	if len(b5) != 5 {
		t.FailNow()
	}
	// check if reversed
	for k, v := range b5 {
		if v.Id+k != 5 {
			t.FailNow()
		}
	}
	x2 := PlacesList{}
	x2 = append(x2, &place.Place{Id: 4}, &place.Place{Id: 5})
	b52 := x2.BottomFive()
	//check if reversed
	for k, v := range b52 {
		if v.Id+k != 5 {
			t.FailNow()
		}
	}
}
