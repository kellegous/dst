package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/ungerik/go-cairo"

	"github.com/kellegous/tzimg/pkg"
	"github.com/kellegous/tzimg/pkg/geo"
	"github.com/kellegous/tzimg/pkg/themes"
	"github.com/kellegous/tzimg/pkg/tz"
)

type Flags struct {
	GeoJSONFile string
	DstFile     string
	ThemesFile  string
	Seed        pkg.Seed
}

func (f *Flags) Register(fs *flag.FlagSet) {
	fs.StringVar(
		&f.GeoJSONFile,
		"geojson-file",
		"combined-with-oceans.json",
		"the geojson file for all timezones")
	fs.StringVar(
		&f.DstFile,
		"dst",
		"tzimg.png",
		"where to write the output image")
	fs.StringVar(
		&f.ThemesFile,
		"themes-file",
		"themes.bin",
		"the file containing all the color themes")
	fs.Var(
		&f.Seed,
		"seed",
		"the PRNG seed")
}

func GetBoundingRect(
	zones []*tz.Timezone,
) (*geo.Location, *geo.Location) {
	min := &geo.Location{
		Lat: math.MaxFloat64,
		Lng: math.MaxFloat64,
	}
	max := &geo.Location{
		Lat: -math.MaxFloat64,
		Lng: -math.MaxFloat64,
	}

	for _, zone := range zones {
		for _, poly := range zone.Polygons {
			for _, ring := range poly.Rings {
				for _, loc := range ring {
					lng, lat := loc.Lng, loc.Lat
					min.Lng = math.Min(lng, min.Lng)
					min.Lat = math.Min(lat, min.Lat)
					max.Lng = math.Max(lng, max.Lng)
					max.Lat = math.Max(lat, max.Lat)
				}
			}
		}
	}

	return min, max
}

func RenderTo(
	dst string,
	w float64,
	h float64,
	polys tz.Polygons,
) error {
	s := cairo.NewPDFSurface(dst, w, h, cairo.PDF_VERSION_1_5)

	p := geo.NewMercator(w, h)

	fmt.Printf(
		"%#v %#v\n",
		p.ToPoint(&geo.Location{Lng: -180, Lat: -90}),
		p.ToPoint(&geo.Location{Lng: 180, Lat: 90}),
	)

	s.SetSourceRGB(0.6, 0.6, 0.6)
	for _, poly := range polys {
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
		s.Fill()
	}

	s.Finish()
	return nil
}

func RenderPNG(
	dst string,
	w float64,
	h float64,
	tx geo.Point,
	zones []*tz.Timezone,
	colors []themes.Color,
) error {
	now := time.Date(2020, time.February, 12, 0, 0, 0, 0, time.Local)
	size := math.Max(w, h)

	s := cairo.NewSurface(cairo.FORMAT_ARGB32, int(size), int(size))
	p := geo.NewMercator(size, size)

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
			s.SetSourceRGB(1, 1, 1)
			s.SetLineWidth(0.5)
			s.Stroke()
		}
	}

	s.SetSourceRGBA(1, 1, 1, 0.3)
	s.Translate(tx.X, tx.Y)
	s.Rectangle(0, 0, w, h)
	s.Fill()

	if status := s.WriteToPNG(dst); status != cairo.STATUS_SUCCESS {
		return fmt.Errorf("unable to write png: %d", status)
	}

	return nil
}

func RenderAll(
	dst string,
	w float64,
	h float64,
	zones []*tz.Timezone,
	colors []themes.Color,
) error {
	now := time.Date(2020, time.February, 12, 0, 0, 0, 0, time.Local)

	s := cairo.NewPDFSurface(dst, w, h, cairo.PDF_VERSION_1_5)

	p := geo.NewMercator(w, h)

	for _, zone := range zones {
		loc, err := time.LoadLocation(zone.ID)
		if err != nil {
			fmt.Printf("ignoring %s\n", zone.ID)
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
			s.SetSourceRGB(1, 1, 1)
			s.SetLineWidth(0.5)
			s.Stroke()
		}
	}

	s.Finish()
	return nil
}

func FindTimezone(zones []*tz.Timezone, id string) *tz.Timezone {
	for _, zone := range zones {
		if zone.ID == id {
			return zone
		}
	}
	return nil
}

func ExpandColors(colors []themes.Color) []themes.Color {
	var nc []themes.Color

	nc = append(nc, colors...)

	white := themes.Color(0xffffff)
	for i := len(colors) - 1; i >= 0; i-- {
		nc = append(nc, colors[i].Lerp(white, 0.5))
	}

	black := themes.Color(0)
	for _, color := range colors {
		nc = append(nc, color.Lerp(black, 0.5))
	}

	return nc
}

func main() {
	var flags Flags
	flags.Register(flag.CommandLine)
	flag.Parse()

	fmt.Printf("seed = %s\n", flags.Seed.String())

	zones, err := tz.ReadFrom(flags.GeoJSONFile)
	if err != nil {
		log.Panic(err)
	}

	ts, err := themes.Open(flags.ThemesFile)
	if err != nil {
		log.Panic(err)
	}
	defer ts.Close()

	theme, err := ts.Pick(flags.Seed.Rand())
	if err != nil {
		log.Panic(err)
	}

	if err := RenderPNG(
		flags.DstFile,
		1600,
		600,
		geo.Point{X: 0, Y: 400},
		zones,
		ExpandColors(theme.Colors),
	); err != nil {
		log.Panic(err)
	}
}
