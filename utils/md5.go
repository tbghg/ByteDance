package utils

import (
	"ByteDance/pkg/common"
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	b := []byte(str)
	s := []byte(common.MD5Salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
