package authority_info

import (
	"XiaoBaoSecurity/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 从jwt中解析登录的信息
type JWTAuthorityInfoMiddleware struct {
	utils utils.JWTUtils
}

// 初始化
func NewJWTAuthorityInfoMiddleware(utils utils.JWTUtils) AuthorityInfoMiddleware {
	return &JWTAuthorityInfoMiddleware{
		utils: utils,
	}
}

// 解析登录信息，解析后放到ctx中
func (j *JWTAuthorityInfoMiddleware) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		fmt.Printf(path)
		//这里是不是可以考虑，尝试获取jwt，也不判断是否需要jwt，拿不到的话，就是没有嘛，没有就不管，
		//后面权限校验的时候会卡掉的
		token := context.GetHeader("x-jwt-token")
		if token != "" {
			////可以去解析
			info, err := j.utils.AnalysisJWT(context, token)
			if err != nil {
				return
			}
			context.Set("info", info)
		}
	}
}
