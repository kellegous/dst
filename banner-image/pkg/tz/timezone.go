package tz

import (
	"encoding/json"
	"os"
)

type Timezone struct {
	ID       string
	Polygons Polygons
}

func (t *Timezone) UnmarshalJSON(b []byte) error {
	var s struct {
		Properties struct {
			ID string `json:"tzid"`
		}
		Geometry Polygons `json:"geometry"`
	}

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t.ID = s.Properties.ID
	t.Polygons = s.Geometry
	return nil
}

func ReadFrom(src string) ([]*Timezone, error) {
	r, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var data struct {
		Zones []*Timezone `json:"features"`
	}

	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, err
	}

	return data.Zones, nil
}
