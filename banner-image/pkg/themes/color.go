package themes

import "fmt"

// Color ...
type Color uint32

// R ...
func (c Color) R() float64 {
	r := (uint32(c) >> 16) & 0xff
	return float64(r) / 255.0
}

// G ...
func (c Color) G() float64 {
	g := (uint32(c) >> 8) & 0xff
	return float64(g) / 255.0
}

// B ...
func (c Color) B() float64 {
	b := uint32(c) & 0xff
	return float64(b) / 255.0
}

func (c Color) String() string {
	return fmt.Sprintf(
		"#%06x (%0.2f, %0.2f, %0.2f)",
		uint32(c),
		c.R(),
		c.G(),
		c.B())
}

// Luminance ...
func (c Color) Luminance() float64 {
	return 0.2126*c.R() + 0.7152*c.G() + 0.0722*c.B()
}

func (c Color) Lerp(b Color, f float64) Color {
	cr, cg, cb := c.R(), c.G(), c.B()
	br, bg, bb := b.R(), b.G(), b.B()
	return ColorFromRGB(
		cr+(br-cr)*f,
		cg+(bg-cg)*f,
		cb+(bb-cb)*f,
	)
}

func ColorFromRGB(r, g, b float64) Color {
	return Color(int32(255*r)<<16 | int32(255*g)<<8 | int32(255*b))
}
