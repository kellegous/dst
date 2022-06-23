package svg

func Doc(box *ViewBox) *Node {
	return NewNode("svg").
		AddAttribute("xmlns", "http://www.w3.org/2000/svg").
		AddAttribute("viewBox", box.String())
}

func Rect(x, y, w, h float64) *LeafNode {
	return NewLeafNode("rect").
		AddAttribute("x", x).
		AddAttribute("y", y).
		AddAttribute("width", w).
		AddAttribute("height", h)
}

func Line(x1, y1, x2, y2 float64) *LeafNode {
	return NewLeafNode("line").
		AddAttribute("x1", x1).
		AddAttribute("y1", y1).
		AddAttribute("x2", x2).
		AddAttribute("y2", y2)
}

func Text(x, y float64, v string) *Node {
	t := NewTextNode(v)
	return NewNode("text").
		AddAttribute("x", x).
		AddAttribute("y", y).
		Append(t)
}

func TextLines(x, y float64, lines []string, dy string) *Node {
	text := NewNode("text").
		AddAttribute("x", x).
		AddAttribute("y", y)
	for _, line := range lines {
		NewNode("tspan").
			AddAttribute("x", x).
			AddAttribute("dy", dy).
			Append(NewTextNode(line)).
			AppendTo(text)
	}
	return text
}

func Path() *LeafNode {
	return NewLeafNode("path")
}

func Circle(cx, cy float64, r float64) *LeafNode {
	return NewLeafNode("circle").
		AddAttribute("cx", cx).
		AddAttribute("cy", cy).
		AddAttribute("r", r)
}
