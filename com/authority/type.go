package authority

import (
	"XiaoBaoSecurity/domian"
	"github.com/gin-gonic/gin"
)

// 权限验证部分
type AuthorityHandler interface {
	//去设置权限
	SetAuthority(fn func(ctx *gin.Context) (domian.AuthorityUserInfo, error)) gin.HandlerFunc
	//静态路由权限校验
	CheckAuthority(fn func(ctx *gin.Context)) gin.HandlerFunc
}
