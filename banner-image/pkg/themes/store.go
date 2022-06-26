package themes

import (
	"encoding/binary"
	"math/rand"

	"golang.org/x/exp/mmap"
)

// Store ...
type Store struct {
	r *mmap.ReaderAt
}

// Close ...
func (s *Store) Close() error {
	return s.r.Close()
}

// Get ...
func (s *Store) Get(idx int) (*Theme, error) {
	var buf [20]byte
	if _, err := s.r.ReadAt(buf[:], int64(idx*5*4)); err != nil {
		return nil, err
	}

	clrs := make([]Color, 5)
	for i := 0; i < 5; i++ {
		clrs[i] = Color(binary.BigEndian.Uint32(buf[i*4:]))
	}

	return &Theme{Index: idx, Colors: clrs}, nil
}

// Pick ...
func (s *Store) Pick(rng *rand.Rand) (*Theme, error) {
	return s.Get(rng.Intn(s.Len()))
}

// Len ...
func (s *Store) Len() int {
	return s.r.Len() / 20
}

// Open ...
func Open(src string) (*Store, error) {
	r, err := mmap.Open(src)
	if err != nil {
		return nil, err
	}

	return &Store{
		r: r,
	}, nil
}
