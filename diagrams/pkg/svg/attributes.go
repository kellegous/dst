package svg

import (
	"fmt"
)

type Attributes map[string]string

func (a Attributes) Add(key string, val any) {
	switch t := val.(type) {
	case int32, int64, int:
		a[key] = fmt.Sprintf("%d", t)
	case float32, float64:
		a[key] = fmt.Sprintf("%0.6f", t)
	case string:
		a[key] = t
	}
}
