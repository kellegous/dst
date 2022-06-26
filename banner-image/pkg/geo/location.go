package geo

import "encoding/json"

type Location struct {
	Lat, Lng float64
}

func (l *Location) UnmarshalJSON(b []byte) error {
	var loc []float64

	if err := json.Unmarshal(b, &loc); err != nil {
		return err
	}

	l.Lng = loc[0]
	l.Lat = loc[1]
	return nil
}
