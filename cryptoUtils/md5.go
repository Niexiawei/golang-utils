package cryptoUtils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return hex.EncodeToString(m.Sum(nil))
}
