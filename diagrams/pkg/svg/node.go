package svg

import (
	"fmt"
	"html/template"
	"io"
)

type Node struct {
	Name       string
	Attributes Attributes
	Children   []Serializer
}

func (n *Node) Append(c Serializer) *Node {
	n.Children = append(n.Children, c)
	return n
}

func (n *Node) AppendTo(p *Node) *Node {
	p.Append(n)
	return n
}

func (n *Node) AddAttribute(key string, val any) *Node {
	n.Attributes.Add(key, val)
	return n
}

func (n *Node) Serialize(w io.Writer) error {
	if len(n.Attributes) == 0 {
		if _, err := fmt.Fprintf(w, "<%s>", n.Name); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprintf(w, "<%s ", n.Name); err != nil {
			return err
		}

		for key, val := range n.Attributes {
			if _, err := fmt.Fprintf(w, `%s="%s" `, key, template.HTMLEscapeString(val)); err != nil {
				return err
			}
		}

		if _, err := fmt.Fprint(w, ">"); err != nil {
			return err
		}
	}

	for _, child := range n.Children {
		if err := child.Serialize(w); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(w, "</%s>", n.Name); err != nil {
		return err
	}

	return nil
}

func NewNode(name string) *Node {
	return &Node{
		Name:       name,
		Attributes: map[string]string{},
	}
}

type LeafNode struct {
	Name       string
	Attributes Attributes
}

func (n *LeafNode) AddAttribute(key string, val any) *LeafNode {
	n.Attributes.Add(key, val)
	return n
}

func (n *LeafNode) AppendTo(p *Node) *LeafNode {
	p.Append(n)
	return n
}

func (n *LeafNode) Serialize(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "<%s ", n.Name); err != nil {
		return err
	}

	for key, val := range n.Attributes {
		if _, err := fmt.Fprintf(w, `%s="%s" `, key, template.HTMLEscapeString(val)); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "/>"); err != nil {
		return err
	}

	return nil
}

func NewLeafNode(name string) *LeafNode {
	return &LeafNode{
		Name:       name,
		Attributes: map[string]string{},
	}
}

type TextNode string

func NewTextNode(v string) TextNode {
	return TextNode(template.HTMLEscapeString(v))
}

func (n TextNode) Serialize(w io.Writer) error {
	if _, err := fmt.Fprint(w, n); err != nil {
		return err
	}
	return nil
}
