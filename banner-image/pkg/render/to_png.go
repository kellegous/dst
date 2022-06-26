package render

import (
	"fmt"
	"math"
	"time"

	"github.com/kellegous/tzimg/pkg/geo"
	"github.com/kellegous/tzimg/pkg/themes"
	"github.com/kellegous/tzimg/pkg/tz"
	"github.com/ungerik/go-cairo"
)

func expandColors(colors []themes.Color) []themes.Color {
	var nc []themes.Color

	nc = append(nc, colors...)

	white := themes.Color(0xffffff)
	for _, color := range colors {
		nc = append(nc, color.Lerp(white, 0.5))
	}
	// white := themes.Color(0xffffff)
	// for i := len(colors) - 1; i >= 0; i-- {
	// 	nc = append(nc, colors[i].Lerp(white, 0.5))
	// }

	// black := themes.Color(0)
	// for _, color := range colors {
	// 	nc = append(nc, color.Lerp(black, 0.5))
	// }

	// return nc
	return nc
}

func ToPNG(
	dst string,
	w float64,
	h float64,
	tx geo.Point,
	now time.Time,
	zones []*tz.Timezone,
	theme *themes.Theme,
) error {
	size := math.Max(w, h)
	colors := expandColors(theme.Colors)

	s := cairo.NewSurface(cairo.FORMAT_ARGB32, int(w), int(h))
	p := geo.NewMercator(size, size)

	s.Translate(tx.X, tx.Y)

	for _, zone := range zones {
		loc, err := time.LoadLocation(zone.ID)
		if err != nil {
			continue
		}

		h := now.In(loc).Hour()
		c := colors[h%len(colors)]

		for _, poly := range zone.Polygons {
			s.NewPath()
			for _, ring := range poly.Rings {
				pt := p.ToPoint(ring[0])
				s.MoveTo(pt.X, pt.Y)
				for i, n := 1, len(ring); i < n; i++ {
					pt := p.ToPoint(ring[i])
					s.LineTo(pt.X, pt.Y)
				}
				s.ClosePath()
			}
			s.SetSourceRGB(c.R(), c.G(), c.B())
			s.FillPreserve()
			s.SetSourceRGB(0.2, 0.2, 0.2)
			s.SetLineWidth(0.5)
			s.Stroke()
		}
	}

	if status := s.WriteToPNG(dst); status != cairo.STATUS_SUCCESS {
		return fmt.Errorf("unable to write png: %d", status)
	}

	return nil
}
