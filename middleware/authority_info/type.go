package authority_info

import "github.com/gin-gonic/gin"

// 这是一个解析登录信息的middleware
type AuthorityInfoMiddleware interface {
	//在这一步解析用户信息,以及做一些续约的事情
	Build() gin.HandlerFunc //构造
}
