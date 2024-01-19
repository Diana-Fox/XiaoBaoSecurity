package main

import (
	"XiaoBaoSecurity/ioc"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	handle := ioc.InitUserHandle("我的密钥，以后可以读配置文件")
	handle.Register(server)
	err := server.Run(":18080")
	if err != nil {
		panic("启动失败")
	}
}
