package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/os/gctx"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
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

// JWTAuth 认证中间件
func JWTAuth(r *ghttp.Request) {
	// 从 Header 的 Authorization 字段获取 Token
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "未提供 token"))
		return
	}

	// 删除 Bearer 前缀
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// 绑定并校验 token
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecretKey), nil
	})
	if err != nil || !token.Valid {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "不合法的 token"))
		return
	}

	// 在 Context 里存储用户名
	r.SetCtxVar(CtxUsername, claims.Username)
	r.SetCtxVar(CtxUserRole, claims.UserRole)
	r.Middleware.Next()
}
