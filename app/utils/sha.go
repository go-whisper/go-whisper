package utils

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	s := h.Sum(nil)
	return fmt.Sprintf("%x", s)
}
