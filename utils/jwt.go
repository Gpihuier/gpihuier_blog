package utils

import (
	"errors"
	"time"

	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/global"

	"github.com/golang-jwt/jwt/v4"
)

type JwtSecret struct {
	signingKey []byte
}

func NewJwtSecret() *JwtSecret {
	return &JwtSecret{
		signingKey: []byte(global.CONFIG.Jwt.SigningKey),
	}
}

func (j *JwtSecret) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	timestamp := time.Now().Unix()
	return request.CustomClaims{
		BaseClaims: baseClaims,                   // 基础数据
		BufferTime: global.CONFIG.Jwt.BufferTime, // 缓存时间一天 // TODO 快要过期时重新获取新的token
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Unix(timestamp-60, 0)),                            // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Unix(timestamp+global.CONFIG.Jwt.ExpiresTime, 0)), // 过期时间 7天
			Issuer:    global.CONFIG.Jwt.Issuer,                                                  // 签发者
		},
	}
}

// CreateToken 创建一个token
func (j *JwtSecret) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signingKey)
}

func (j *JwtSecret) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})
	if err != nil { // TODO 增加多错误判断
		return nil, err
	}
	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
