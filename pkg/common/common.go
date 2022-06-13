package common

import (
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
