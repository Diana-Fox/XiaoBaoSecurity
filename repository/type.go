package repository

import (
	"XiaoBaoSecurity/repository/dao"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, email string, password string) error
	FindByEmail(ctx context.Context, email string) (dao.User, error)
	//查询到这个用户的资源信息
	FindUserAuthority(ctx context.Context, uid int64) ([]dao.Resource, error)
}
