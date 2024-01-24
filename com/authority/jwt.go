package authority

import (
	"XiaoBaoSecurity/domian"
	"XiaoBaoSecurity/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type JWTAuthority struct {
	jwt utils.JWTUtils
}

func NewAuthority(jwt utils.JWTUtils) AuthorityHandler {
	return &JWTAuthority{jwt: jwt}
}

// 提供给登录接口使用，用于写登录相关内容
func (j *JWTAuthority) SetAuthority(fn func(ctx *gin.Context) (domian.AuthorityUserInfo, error)) gin.HandlerFunc {
	return func(context *gin.Context) {
		info, err := fn(context)
		if err != nil {
			context.String(401, err.Error())
			return
		}
		//去生成jwt并存到请求头上
		err = j.jwt.GenerateJWT(context, info)
		if err != nil {
			context.String(403, "权限设置异常")
			return
		}
		context.String(200, "登录成功")
	}
}

func (j *JWTAuthority) CheckAuthority(fn func(ctx *gin.Context)) gin.HandlerFunc {
	return func(context *gin.Context) {
		info, err := j.getInfo(context)
		if err != nil {
			//todo 错误了，直接结束,可以，返回值不是重点，所以暂时这样就行了
			context.String(401, err.Error())
			return
		}
		path := context.FullPath()
		urlMap := info.UrlMap
		_, ok := urlMap[path]
		if ok {
			//说明在map中命中了,可以执行了
			fmt.Printf("在map")
			fn(context)
			return
		}
		context.String(401, errors.New("没有当前路径权限").Error())
	}
}

func (j *JWTAuthority) getInfo(ctx *gin.Context) (domian.AuthorityUserInfo, error) {
	value, exists := ctx.Get("info")
	if !exists {
		//没解析到
		return domian.AuthorityUserInfo{}, errors.New("用户信息不存在")
	}
	info, ok := value.(domian.AuthorityUserInfo)
	if !ok {
		//解析的不对
		return domian.AuthorityUserInfo{}, errors.New("用户信息类型错误")
	}
	return info, nil
}
