package pkg

import (
	crand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
)

type Seed int64

func (s *Seed) Set(v string) error {
	b, err := hex.DecodeString(v)
	if err != nil {
		return err
	}

	if len(b) != 8 {
		return errors.New("seed must be 8 characters")
	}

	*s = Seed(int64(binary.BigEndian.Uint64(b)))
	return nil
}

func (s *Seed) String() string {
	s.init()
	return fmt.Sprintf("%016x", uint64(*s))
}

func (s *Seed) Rand() *rand.Rand {
	s.init()
	return rand.New(rand.NewSource(int64(*s)))
}

func (s Seed) Seed() int64 {
	s.init()
	return int64(s)
}

func (s *Seed) init() {
	if *s != 0 {
		return
	}

	var buf [8]byte
	crand.Read(buf[:])
	*s = Seed(binary.BigEndian.Uint64(buf[:]))
}
