package tz

import (
	"encoding/json"
)

type Point struct {
	X, Y float64
}

func (p *Point) UnmarshalJSON(b []byte) error {
	var pt []float64
	if err := json.Unmarshal(b, &pt); err != nil {
		return err
	}
	p.X = pt[0]
	p.Y = pt[1]

	return nil
}
