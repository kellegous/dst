package geo

import (
	"math"
)

const quarterPi = math.Pi / 4

type Mercator struct {
	w, h float64
}

func NewMercator(w, h float64) *Mercator {
	return &Mercator{
		w: w,
		h: h,
	}
}

func toPixelY(lat float64, h float64) float64 {
	lat = math.Max(math.Min(lat, 89.99), -89.99)
	lr := lat * math.Pi / 180
	mn := math.Log(math.Tan(quarterPi + lr/2))
	y := (h / 2) - (h * mn / (2 * math.Pi))
	return y
}

func (m *Mercator) ToPoint(loc *Location) Point {
	x := (loc.Lng + 180) * (m.w / 360)
	y := toPixelY(loc.Lat, m.h)
	return Point{X: x, Y: y}
}
