package service

import (
	"XiaoBaoSecurity/domian"
	"context"
)

// 用户
type UserService interface {
	SignUp(ctx context.Context, email string, password string) error
	LoginByEmail(ctx context.Context, email string, password string) (domian.AuthorityUserInfo, error)
}
