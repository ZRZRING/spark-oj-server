package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GenToken 生成 JWT token
func GenToken(username string, userRole string) (token string, err error) {
	// 创建 JWT 载荷
	claims := &JWTClaims{
		Username: username,
		UserRole: userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  []string{},                                         // 受众
			Issuer:    "SparkOJ",                                          // 签发人
			Subject:   "JWTAuth",                                          // 主题
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 有效时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			ID:        "main",                                             // JWT ID
		},
	}

	// 生成 JWT token
	JWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = JWT.SignedString([]byte(JwtSecretKey))
	if err != nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, "生成 token 失败")
		return
	}

	return
}
