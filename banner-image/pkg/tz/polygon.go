package tz

import (
	"encoding/json"

	"github.com/kellegous/tzimg/pkg/geo"
)

type Polygon struct {
	Rings [][]*geo.Location
}

func (p *Polygon) UnmarshalJSON(b []byte) error {
	var locs [][]*geo.Location
	if err := json.Unmarshal(b, &locs); err != nil {
		return err
	}
	p.Rings = locs
	return nil
}
