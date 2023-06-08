package cryptohelpers

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(input string) string {
	h := md5.New()
	defer h.Reset()
	h.Write([]byte(input))

	return hex.EncodeToString(h.Sum(nil))
}
