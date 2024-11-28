package main

import (
	"crypto/sha1"
	"encoding/binary"
)

func hashKey(key string) uint32 {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	hashBytes := hasher.Sum(nil)
	return binary.BigEndian.Uint32(hashBytes[:4]) // Take the first 4 bytes
}
