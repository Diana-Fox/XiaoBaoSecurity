package repository

import (
	"XiaoBaoSecurity/repository/dao"
	"context"
	"database/sql"
)

type DefaultUserRepository struct {
	ud  dao.UserDao
	rd  dao.UserRoleDao
	rrd dao.RoleResourceDao
	rsd dao.ResourceDao
}

func (d *DefaultUserRepository) FindUserAuthority(ctx context.Context, uid int64) ([]dao.Resource, error) {
	rIds, err := d.rd.FindRIdsByUId(ctx, uid)
	if err != nil {
		return nil, err
	}
	//拿到rIds后
	resIds, err := d.rrd.FindResourceByRoles(ctx, rIds)
	if err != nil {
		return nil, err
	}
	err, resList := d.rsd.FindByIds(ctx, resIds)
	return resList, err
}

func NewDefaultUserRepository(ud dao.UserDao, rd dao.UserRoleDao,
	rrd dao.RoleResourceDao, rsd dao.ResourceDao) UserRepository {
	return &DefaultUserRepository{
		ud:  ud,
		rd:  rd,
		rrd: rrd,
		rsd: rsd,
	}
}
func (d *DefaultUserRepository) Create(ctx context.Context, email string, password string) error {
	return d.ud.Insert(ctx, dao.User{
		Email: sql.NullString{
			String: email,
			Valid:  email != "",
		},
		Password: password,
	})
}

func (d *DefaultUserRepository) FindByEmail(ctx context.Context, email string) (dao.User, error) {
	return d.ud.FindByEmail(ctx, email)
}
