package utils

import (
	"ByteDance/pkg/common"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

/*
JWT使用
*/

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 这里额外记录username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

// mySecret 密钥
var mySecret = []byte(common.MySecret)

// GenToken 生成 Token
func GenToken(id int) (string, error) {
	c := MyClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    common.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(common.TokenExpirationTime)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                 // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                 // 生效时间
		}}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}
