package list

import (
	"github.com/yaronsumel/geo-coding-test/src/place"
	"sort"
)

// PlacesList
// Holds Places
type PlacesList []*place.Place

// Len implements Len() for sort functions
func (s PlacesList) Len() int {
	return len(s)
}

// Swap implements Swap(i, j int) for sort functions
func (s PlacesList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less implements Less(i, j int) for sort functions
func (s PlacesList) Less(i, j int) bool {
	return s[i].Distance < s[j].Distance
}

// Append
// grow list and sort items
// if reached more than ten items
// delete the middle item after sort
// so we can keep top 5 and bottom 5
func (s *PlacesList) Append(p *place.Place) {
	*s = append(*s, p)
	// sort current items
	// will use insertionSort
	sort.Stable(s)
	// if sorted list is bigger than 10
	// delete the middle one
	// so we still got the topFive and bottomFive
	if len(*s) > 10 {
		*s = s.DeleteI(5)
	}
}

// DeleteI
// delete item i from PlacesList and return new list
func (s PlacesList) DeleteI(i int) PlacesList {
	return append(s[:i], s[i+1:]...)
}

// Top
// return top5 PlacesList
func (s PlacesList) TopFive() PlacesList {
	l := len(s)
	if l < 5 {
		return s[:l]
	}
	return s[:5]
}

// Bottom
// return bottom5 placesList
func (s PlacesList) BottomFive() PlacesList {
	l := len(s)
	// reverse list
	// really cool trick
	// you should look at https://github.com/golang/go/wiki/SliceTricks
	for i := l/2 - 1; i >= 0; i-- {
		opp := l - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
	if l < 5 {
		return s[:l]
	}
	return s[:5]
}
