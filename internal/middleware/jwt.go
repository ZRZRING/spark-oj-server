package middleware

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang-jwt/jwt/v5"
)

// JwtSecretKey 鉴权密钥
const (
	CtxUsername  gctx.StrKey = "username"
	CtxUserRole  gctx.StrKey = "role"
	JwtSecretKey string      = "b47qV7nDxtQpmURfBpx7"
)

// JWTClaims 原始的 JWT 载荷
type JWTClaims struct {
	Username string `json:"username"`
	UserRole string `json:"user_role"`
	jwt.RegisteredClaims
}
