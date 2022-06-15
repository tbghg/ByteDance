package utils

import (
	"ByteDance/pkg/common"
	"crypto/md5"
	"encoding/hex"
	"github.com/dlclark/regexp2"
)

func Md5(str string) string {
	b := []byte(str)
	s := []byte(common.MD5Salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// MatchStr 密码强度检测
func MatchStr(str string) bool {
	expr := `^(?![0-9a-zA-Z]+$)(?![a-zA-Z!@#$%^&*]+$)(?![0-9!@#$%^&*]+$)[0-9A-Za-z!@#$%^&*]{8,32}$`
	reg, _ := regexp2.Compile(expr, 0)
	m, _ := reg.FindStringMatch(str)
	if m != nil {
		return true
	}
	return false
}
