package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

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
