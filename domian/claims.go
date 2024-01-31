package domian

import "github.com/golang-jwt/jwt/v5"

type UserClaims[T any] struct {
	jwt.RegisteredClaims
	UserInfo T //存权限信息
}
