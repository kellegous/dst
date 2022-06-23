package svg

import "io"

type Serializer interface {
	Serialize(w io.Writer) error
}
