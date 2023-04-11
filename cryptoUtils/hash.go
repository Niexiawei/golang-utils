package cryptoUtils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func HashSHA256(str string) string {
	s := sha256.New()
	io.WriteString(s, str)
	return hex.EncodeToString(s.Sum(nil))
}
