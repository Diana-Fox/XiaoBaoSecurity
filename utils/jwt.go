package utils

import (
	"XiaoBaoSecurity/domian"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// JWTUtils 改成泛型
type JWTUtils[T any] interface {
	GenerateJWT(ctx *gin.Context, userInfo T) error        //生成jwt
	AnalysisJWT(ctx *gin.Context, token string) (T, error) //解析jwt
}

type jwtUtils struct {
	secretKey string
}

// 密钥一般不会变，所有固定一下
func NewJWTUtils(secretKey string) JWTUtils[domian.AuthorityUserInfo] {
	return &jwtUtils{
		secretKey: secretKey,
	}
}
func (j jwtUtils) GenerateJWT(ctx *gin.Context, userInfo domian.AuthorityUserInfo) error {
	var claims = domian.UserClaims[domian.AuthorityUserInfo]{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)), //
		},
		UserInfo: userInfo,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", signedString)
	return nil
}

func (j jwtUtils) AnalysisJWT(ctx *gin.Context, token string) (domian.AuthorityUserInfo, error) {
	clamis := &domian.UserClaims[domian.AuthorityUserInfo]{}
	parseClaims, err := jwt.ParseWithClaims(token, clamis, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil || !parseClaims.Valid {
		//这里细致化的话，可以考虑定义异常，让上层获取
		return domian.AuthorityUserInfo{}, errors.New("权限验证失败")
	}
	now := time.Now()
	if clamis.ExpiresAt.Sub(now) < time.Second*50 { //刷新
		//10S一续约
		//续约，生成新token
		clamis.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
		signedString, err := parseClaims.SignedString([]byte(j.secretKey))
		if err == nil {
			//顺利续约，其实没续约成功不影响，只要下一次没过期
			ctx.Header("x-jwt-token", signedString)
		}
	}
	return clamis.UserInfo, nil
}
