package place

import "testing"

// TestGreatCircleDistance
func TestCalcGreatCircleDistance(t *testing.T) {
	p1 := Place{Lat: 120, Lon: 120}
	p2 := Place{Lat: 120, Lon: 120}
	p1.CalcGreatCircleDistance(&p2)
	if p1.Distance != 0 {
		t.FailNow()
	}
}
