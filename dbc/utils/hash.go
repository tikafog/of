package utils

import (
	"encoding/binary"
	"sort"

	"github.com/cespare/xxhash"
	"github.com/google/uuid"
)

type Uint64Slice []uint64

func (x Uint64Slice) Len() int           { return len(x) }
func (x Uint64Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x Uint64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func HashString(hashes ...string) []byte {
	hashed := make(Uint64Slice, len(hashes))
	for i := range hashes {
		hashed[i] = xxhash.Sum64String(hashes[i])
	}
	sort.Sort(hashed)
	return Int64ToBytes(hashed...)
}

func Hash(ids ...uuid.UUID) []byte {
	hashed := make(Uint64Slice, len(ids))
	for i := range ids {
		hashed[i] = xxhash.Sum64(ids[i][:])
	}
	sort.Sort(hashed)
	return Int64ToBytes(hashed...)
}

func Int64ToBytes(ints ...uint64) []byte {
	var buf = make([]byte, 8*len(ints))
	for i := range ints {
		binary.BigEndian.PutUint64(buf[i*8:i*8+8], ints[i])
	}
	return buf
}

func BytesToInt64(buf []byte) []uint64 {
	size := len(buf)
	if size == 0 {
		return nil
	}
	ints := make([]uint64, size/8)
	for i := 0; i < size/8; i++ {
		ints[i] = binary.BigEndian.Uint64(buf[i*8 : i*8+8])
	}
	return ints
}
