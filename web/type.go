package web

import (
	"XiaoBaoSecurity/domian"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	//注册路由
	Register(server *gin.Engine)
	//登录
	Login(ctx *gin.Context) (domian.AuthorityUserInfo, error)
}
