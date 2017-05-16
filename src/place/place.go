package place

import "math"

const EARTH_RADIUS = 6371

// Place
type Place struct {
	Id       int
	Lat      float64
	Lon      float64
	Distance float64
	Index    int
}

// GreatCircleDistance calculate distance between two Places
// http://www.movable-type.co.uk/scripts/latlong.html
func (p *Place) CalcGreatCircleDistance(p2 *Place) {
	dLat, dLon := (p2.Lat-p.Lat)*(math.Pi/180.0), (p2.Lon-p.Lon)*(math.Pi/180.0)
	lat1, lat2 := p.Lat*(math.Pi/180.0), p2.Lat*(math.Pi/180.0)
	a1, a2 := math.Sin(dLat/2)*math.Sin(dLat/2), math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	a := a1 + a2
	d := EARTH_RADIUS * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	p.Distance = d
}
