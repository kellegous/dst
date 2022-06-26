package main

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Flags struct {
	SrcFile string
	DstFile string
}

func (f *Flags) Register(fs *flag.FlagSet) {
	fs.StringVar(
		&f.SrcFile,
		"src-file",
		"/Users/knorton/Documents/2018/07/color/kuler.json.gz",
		"the kuler source file")
	fs.StringVar(
		&f.DstFile,
		"dst-file",
		"themes-sorted.bin",
		"the file to write to")
}

type Swatch struct {
	Hex string `json:"hex"`
}

type Theme struct {
	Swatches []Swatch `json:"swatches"`
	Rating   struct {
		Count   int     `json:"count"`
		Overall float64 `json:"overall"`
	} `json:"Rating"`
}

func (t *Theme) Marshal(b []byte) error {
	if len(b) < 20 {
		return errors.New("need at least 20 bytes")
	}

	for i, swatch := range t.Swatches {
		c, err := strconv.ParseUint(swatch.Hex, 16, 32)
		if err != nil {
			return err
		}

		binary.BigEndian.PutUint32(b[i*4:], uint32(c))
	}

	return nil
}

type Iter struct {
	s *bufio.Scanner
	io.Closer
}

func (i *Iter) Next() (*Theme, error) {
	if !i.s.Scan() {
		return nil, i.s.Err()
	}

	theme := &Theme{}
	if err := json.Unmarshal(i.s.Bytes(), &theme); err != nil {
		return nil, err
	}

	return theme, nil
}

func OpenSrc(src string) (*Iter, error) {
	r, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	gr, err := gzip.NewReader(r)
	if err != nil {
		r.Close()
		return nil, err
	}

	return &Iter{
		s:      bufio.NewScanner(gr),
		Closer: r,
	}, nil
}

func main() {
	var flags Flags
	flags.Register(flag.CommandLine)
	flag.Parse()

	iter, err := OpenSrc(flags.SrcFile)
	if err != nil {
		log.Panic(err)
	}
	defer iter.Close()

	var themes []*Theme

	for {
		theme, err := iter.Next()
		if err != nil {
			log.Panic(err)
		} else if theme == nil {
			break
		}

		if len(theme.Swatches) != 5 {
			continue
		}

		if theme.Rating.Count == 0 {
			continue
		}

		themes = append(themes, theme)
	}

	sort.Slice(themes, func(i, j int) bool {
		a := themes[i]
		b := themes[j]
		if a.Rating.Overall == b.Rating.Overall {
			return b.Rating.Count < a.Rating.Count
		}
		return b.Rating.Overall < a.Rating.Overall
	})

	w, err := os.Create(flags.DstFile)
	if err != nil {
		log.Panic(err)
	}
	defer w.Close()

	var buf [20]byte
	for _, theme := range themes {
		if err := theme.Marshal(buf[:]); err != nil {
			log.Panic(err)
		}

		if _, err := w.Write(buf[:]); err != nil {
			log.Panic(err)
		}
	}
}
