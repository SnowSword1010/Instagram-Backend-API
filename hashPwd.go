package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func givePwdHash(s string) string {
	// APPLYING SHA1
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	// APPLYING SHA256 ON TOP OF IT
	g := sha256.New()
	g.Write([]byte(sha1_hash))
	sha256_hash := hex.EncodeToString(g.Sum(nil))
	return sha256_hash
}
