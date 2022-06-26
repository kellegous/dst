package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kellegous/tzimg/pkg"
	"github.com/kellegous/tzimg/pkg/geo"
	"github.com/kellegous/tzimg/pkg/render"
	"github.com/kellegous/tzimg/pkg/themes"
	"github.com/kellegous/tzimg/pkg/tz"
)

type Flags struct {
	N           int
	NWorkers    int
	DstDir      string
	ThemesFile  string
	GeoJSONFile string
}

func (f *Flags) Register(fs *flag.FlagSet) {
	fs.IntVar(
		&f.N,
		"n",
		10,
		"the number of images to generate")
	fs.IntVar(
		&f.NWorkers,
		"n-workers",
		25,
		"the number of workers in the worker pool")
	fs.StringVar(
		&f.GeoJSONFile,
		"geojson-file",
		"combined-with-oceans.json",
		"the geojson file for all timezones")
	fs.StringVar(
		&f.DstDir,
		"dst-dir",
		"out",
		"where to write the output images")
	fs.StringVar(
		&f.ThemesFile,
		"themes-file",
		"themes.bin",
		"the file containing all the color themes")
}

func main() {
	var flags Flags
	flags.Register(flag.CommandLine)
	flag.Parse()

	if _, err := os.Stat(flags.DstDir); err != nil {
		if err := os.MkdirAll(flags.DstDir, 0755); err != nil {
			log.Panic(err)
		}
	}

	zones, err := tz.ReadFrom(flags.GeoJSONFile)
	if err != nil {
		log.Panic(err)
	}

	ts, err := themes.Open(flags.ThemesFile)
	if err != nil {
		log.Panic(err)
	}
	defer ts.Close()

	pool := pkg.StartWorkers(flags.NWorkers)
	for i := 0; i < flags.N; i++ {
		pool.Submit(func() {
			var seed pkg.Seed

			theme, err := ts.Pick(seed.Rand())
			if err != nil {
				log.Panic(err)
			}

			if err := render.ToPNG(
				filepath.Join(flags.DstDir, seed.String()+".png"),
				1600,
				600,
				geo.Point{X: 0, Y: -400},
				time.Date(2022, time.February, 12, 0, 0, 0, 0, time.Local),
				zones,
				theme,
			); err != nil {
				log.Panic(err)
			}
		})
	}

	pool.Drain()
}
