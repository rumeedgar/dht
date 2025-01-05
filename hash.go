 package main

import (
    "crypto/sha1"
    "math/big"
)

// hashKey takes a key (string) and returns its hash as a big integer
func hashKey(key string) *big.Int {
    h := sha1.New()
    h.Write([]byte(key))
    return new(big.Int).SetBytes(h.Sum(nil))
}
