package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kellegous/dst/diagrams/pkg/svg"
)

type Path struct {
	cmds []string
}

func (p *Path) MoveTo(x, y float64) *Path {
	p.cmds = append(p.cmds, fmt.Sprintf("M %0.4f %0.4f", x, y))
	return p
}

func (p *Path) LineTo(x, y float64) *Path {
	p.cmds = append(p.cmds, fmt.Sprintf("L %0.4f %0.4f", x, y))
	return p
}

func (p *Path) D() string {
	return strings.Join(p.cmds, " ")
}

type Point struct {
	X, Y float64
}

func RenderTimeline(
	doc *svg.Node,
	w float64,
	h float64,
	yOff float64,
	pos bool,
) map[string][]Point {
	n := 6

	positions := map[string][]Point{}

	zones := []struct {
		Name   string
		Offset int
	}{
		{"EST", 0},
		{"EDT", 1},
		{"UTC", 5},
	}

	fontSize := 14.0
	xOff := 75.0

	dw := (w - xOff) / float64(n)
	lh := h / float64(len(zones)-1)

	for i, zone := range zones {
		y := yOff + lh*float64(i)
		svg.Line(0, y, w, y).
			AddAttribute("stroke", "#333").
			AppendTo(doc)

		svg.Text(2, y+10, zone.Name).
			AddAttribute("font-family", "Helvetica").
			AddAttribute("font-size", fontSize).
			AddAttribute("fill", "#666").
			AddAttribute("dominant-baseline", "hanging").
			AppendTo(doc)
		var points []Point
		for i := 0; i < n; i++ {
			x := xOff + dw*float64(i)
			points = append(points, Point{X: x, Y: y})
			svg.Text(x+8.0, y+10, fmt.Sprintf("%d:00", i+zone.Offset)).
				AddAttribute("font-family", "Helvetica").
				AddAttribute("font-size", fontSize).
				AddAttribute("dominant-baseline", "hanging").
				AddAttribute("fill", "#666").
				AppendTo(doc)
			svg.Line(x, y, x, y+20).
				AddAttribute("stroke", "#666").
				AppendTo(doc)
		}
		positions[zone.Name] = points
	}

	if pos {
		var p Path
		p.MoveTo(0, yOff).
			LineTo(xOff+dw*2, yOff).
			LineTo(xOff+dw*2, yOff+lh).
			LineTo(w, yOff+lh)

		svg.Path().
			AddAttribute("d", p.D()).
			AddAttribute("stroke", "#09f").
			AddAttribute("fill", "none").
			AddAttribute("stroke-width", 6.0).
			AddAttribute("stroke-linecap", "round").
			AppendTo(doc)

		svg.Text(0, yOff-10, "America/New_York").
			AddAttribute("fill", "#09f").
			AddAttribute("font-family", "Helvetica").
			AddAttribute("font-size", fontSize).
			AppendTo(doc)
	} else {
		var p Path
		p.MoveTo(0, yOff+lh).
			LineTo(xOff+dw, yOff+lh).
			LineTo(xOff+dw, yOff).
			LineTo(w, yOff)
		svg.Path().
			AddAttribute("d", p.D()).
			AddAttribute("stroke", "#09f").
			AddAttribute("fill", "none").
			AddAttribute("stroke-width", 6.0).
			AddAttribute("stroke-linecap", "round").
			AppendTo(doc)

		svg.Text(w-10, yOff-10, "America/New_York").
			AddAttribute("fill", "#09f").
			AddAttribute("font-family", "Helvetica").
			AddAttribute("font-size", fontSize).
			AddAttribute("text-anchor", "end").
			AppendTo(doc)
	}

	return positions
}

func RenderNegSummary(w, h float64) *svg.Node {
	yOff := (h - 175)
	doc := svg.Doc(svg.NewViewBox(0, 0, w, h))
	pos := RenderTimeline(
		doc,
		w,
		175-50,
		yOff+25,
		false)

	a := pos["EST"][1]
	svg.Text(a.X+100+10, a.Y-100, "Go, JavaScript, Java, PHP, Python, & Ruby").
		AddAttribute("fill", "#666").
		AddAttribute("font-family", "Helvetica").
		AddAttribute("font-size", 14).
		AppendTo(doc)

	svg.Line(a.X, a.Y, a.X+100, a.Y-100).
		AddAttribute("stroke", "#333").
		AddAttribute("stroke-width", 2).
		AppendTo(doc)

	svg.Circle(a.X, a.Y, 6).
		AddAttribute("stroke", "#333").
		AddAttribute("fill", "#fff").
		AddAttribute("stroke-width", 4).
		AppendTo(doc)

	return doc
}

func RenderPosSummary(w, h float64) *svg.Node {
	yOff := (h - 175)
	doc := svg.Doc(svg.NewViewBox(0, 0, w, h))
	pos := RenderTimeline(
		doc,
		w,
		175-50,
		yOff+25,
		true)

	a := pos["EST"][1]
	svg.Text(a.X+100+10, a.Y-100, "Go").
		AddAttribute("fill", "#666").
		AddAttribute("font-family", "Helvetica").
		AddAttribute("font-size", 14).
		AppendTo(doc)
	svg.Line(a.X, a.Y, a.X+100, a.Y-100).
		AddAttribute("stroke", "#333").
		AddAttribute("stroke-width", 2).
		AppendTo(doc)
	svg.Circle(a.X, a.Y, 6).
		AddAttribute("stroke", "#333").
		AddAttribute("fill", "#fff").
		AddAttribute("stroke-width", 4).
		AppendTo(doc)

	b := pos["EDT"][2]
	svg.Text(
		b.X+100+10,
		b.Y-100,
		"JavaScript, Java, PHP, & Ruby").
		AddAttribute("fill", "#666").
		AddAttribute("font-family", "Helvetica").
		AddAttribute("font-size", 14).
		AppendTo(doc)
	svg.Line(b.X, b.Y, b.X+100, b.Y-100).
		AddAttribute("stroke", "#333").
		AddAttribute("stroke-width", 2).
		AppendTo(doc)
	svg.Circle(b.X, b.Y, 6).
		AddAttribute("stroke", "#333").
		AddAttribute("fill", "#fff").
		AddAttribute("stroke-width", 4).
		AppendTo(doc)

	c := pos["EDT"][1]
	svg.Text(c.X+100+10, c.Y-100, "Python").
		AddAttribute("font-family", "Helvetica").
		AddAttribute("font-size", 14).
		AddAttribute("fill", "#666").
		AppendTo(doc)
	svg.Line(c.X, c.Y, c.X+100, c.Y-100).
		AddAttribute("stroke", "#333").
		AddAttribute("stroke-width", 2).
		AppendTo(doc)
	svg.Circle(c.X, c.Y, 6).
		AddAttribute("stroke", "#333").
		AddAttribute("fill", "#fff").
		AddAttribute("stroke-width", 4).
		AppendTo(doc)

	return doc
}

func RenderPos(w, h float64) *svg.Node {
	doc := svg.Doc(svg.NewViewBox(0, 0, w, h))
	RenderTimeline(
		doc,
		w,
		175-50,
		25,
		true)
	return doc
}

func RenderNeg(w, h float64) *svg.Node {
	doc := svg.Doc(svg.NewViewBox(0, 0, w, h))
	RenderTimeline(
		doc,
		w,
		175-50,
		25,
		false)
	return doc
}

func WriteFile(dst string, n *svg.Node) error {
	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()

	return n.Serialize(w)
}

func main() {
	if err := WriteFile(
		"pos-lg.svg",
		RenderPos(900, 175),
	); err != nil {
		log.Panic(err)
	}

	if err := WriteFile(
		"pos-sm.svg",
		RenderPos(450, 175),
	); err != nil {
		log.Panic(err)
	}

	if err := WriteFile(
		"pos-lg-summary.svg",
		RenderPosSummary(900, 275),
	); err != nil {
		log.Panic(err)
	}
	if err := WriteFile(
		"pos-sm-summary.svg",
		RenderPosSummary(450, 275),
	); err != nil {
		log.Panic(err)
	}

	if err := WriteFile(
		"neg-lg.svg",
		RenderNeg(900, 175),
	); err != nil {
		log.Panic(err)
	}

	if err := WriteFile(
		"neg-sm.svg",
		RenderNeg(450, 175),
	); err != nil {
		log.Panic(err)
	}

	if err := WriteFile(
		"neg-lg-summary.svg",
		RenderNegSummary(900, 275),
	); err != nil {
		log.Panic(err)
	}
	if err := WriteFile(
		"neg-sm-summary.svg",
		RenderNegSummary(450, 275),
	); err != nil {
		log.Panic(err)
	}
}
