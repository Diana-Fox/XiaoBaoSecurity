package domian

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.RegisteredClaims
	UserInfo AuthorityUserInfo //存权限信息
}
