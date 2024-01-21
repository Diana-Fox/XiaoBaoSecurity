package main

import (
	"XiaoBaoSecurity/ioc"
	"XiaoBaoSecurity/middleware/authority_info"
	"XiaoBaoSecurity/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	jwtUtils := utils.NewJWTUtils("我的密钥")
	//注意顺序，解析信息的middleware要在前面
	server.Use(authority_info.NewJWTAuthorityInfoMiddleware(jwtUtils).Build())
	handle := ioc.InitUserHandle(jwtUtils)
	handle.Register(server)
	err := server.Run(":18080")
	if err != nil {
		panic("启动失败")
	}
}
