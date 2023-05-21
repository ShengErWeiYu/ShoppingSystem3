package untils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func ComplexPassword(password string) string {
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		panic(err)
	}
	sum := h.Sum(nil)
	return hex.EncodeToString(sum[:])
}
