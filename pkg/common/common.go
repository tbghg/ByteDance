package common

import (
	"github.com/dlclark/regexp2"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// Response 响应共有响应头
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

//参数检验器
var (
	Validate = validator.New()          // 实例化验证器
	Chinese  = zh.New()                 // 获取中文翻译器
	Uni      = ut.New(Chinese, Chinese) // 设置成中文翻译器
	Trans, _ = Uni.GetTranslator("zh")  // 获取翻译字典
)

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
