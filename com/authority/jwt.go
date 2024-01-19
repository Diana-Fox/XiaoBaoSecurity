package authority

import (
	"XiaoBaoSecurity/domian"
	"XiaoBaoSecurity/utils"
	"errors"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"strings"
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
			//todo 登录失败，回头写点日志
			return
		}
		//去生成jwt并存到请求头上
		err = j.jwt.GenerateJWT(context, info)
		if err != nil {
			//todo 登录失败，回头写点日志
			return
		}
	}
}

func (j *JWTAuthority) CheckStaticAuthority(fn func(ctx *gin.Context)) gin.HandlerFunc {
	return func(context *gin.Context) {
		info, err := j.getInfo(context)
		if err != nil {
			//todo 错误了，直接结束,可以，返回值不是重点，所以暂时这样就行了
			context.String(401, err.Error())
			return
		}
		path := context.Request.URL.Path
		urlMap := info.UrlMap //静态路由，所以只用map
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

func (j *JWTAuthority) CheckDynamicsAuthority(fn func(ctx *gin.Context)) gin.HandlerFunc {
	return func(context *gin.Context) {
		info, err := j.getInfo(context)
		if err != nil {
			//todo 错误了，直接结束,返回值不是重点，所以暂时这样就行了
			context.String(401, err.Error())
			return
		}
		path := context.Request.URL.Path
		pathLevel := len(strings.Split(path, `/`))
		urlList := info.UrlList //动态路由，所以去遍历然后匹配正则
		for _, node := range urlList {
			if node.Level == pathLevel {
				//一样长度，所以值得正则一下
				ok, err := regexp.MustCompile(node.Url, regexp.None).MatchString(path)
				if err != nil {
					//todo 这里后面直接拒绝请求，理论上不会报错
					fmt.Printf("正则报错")
					context.String(401, "信息异常")
					return
				}
				if ok {
					//要是匹配上了，直接放行
					fn(context)
					return
				}
			}
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
