package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/kellegous/tzimg/pkg"
	"github.com/kellegous/tzimg/pkg/geo"
	"github.com/kellegous/tzimg/pkg/render"
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

	if err := render.ToPNG(
		flags.DstFile,
		1600,
		600,
		geo.Point{X: 0, Y: -400},
		time.Date(2022, time.February, 12, 0, 0, 0, 0, time.Local),
		zones,
		theme,
	); err != nil {
		log.Panic(err)
	}
}
