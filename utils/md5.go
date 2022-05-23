package utils

import (
	"ByteDance/config"
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	b := []byte(str)
	s := []byte(config.Salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
