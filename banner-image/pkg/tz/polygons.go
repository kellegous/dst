package tz

import (
	"encoding/json"
)

type Polygons []*Polygon

func (p *Polygons) UnmarshalJSON(b []byte) error {
	var data struct {
		Type        string          `json:"type"`
		Coordinates json.RawMessage `json:"coordinates"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	switch data.Type {
	case "Polygon":
		var poly Polygon
		if err := json.Unmarshal(data.Coordinates, &poly); err != nil {
			return err
		}
		*p = []*Polygon{&poly}
	case "MultiPolygon":
		var polys []*Polygon
		if err := json.Unmarshal(data.Coordinates, &polys); err != nil {
			return err
		}
		*p = polys
	}

	return nil
}
