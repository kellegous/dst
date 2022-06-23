package svg

import "fmt"

type ViewBox struct {
	X float64
	Y float64
	W float64
	H float64
}

func (b *ViewBox) String() string {
	return fmt.Sprintf("%f %f %f %f", b.X, b.Y, b.W, b.H)
}

func NewViewBox(x, y, w, h float64) *ViewBox {
	return &ViewBox{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}
