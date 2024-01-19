package dao

import "context"

// 用户
type UserDao interface {
	Insert(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
}

// 角色
type RoleDao interface {
	Insert(ctx context.Context, role Role) error
	FindAll(ctx context.Context) (Role, error)                  //这里需要按父子关系组织成树
	FindByIds(ctx context.Context, ids []int64) ([]Role, error) //批量查找
}

// 资源
type ResourceDao interface {
	Insert(ctx context.Context, resource Resource) error
	FindAll(ctx context.Context) (error, Resource)                  //这里也是查询所有的，也是前端维护页面需要的
	FindByIds(ctx context.Context, ids []int64) (error, []Resource) //批量查找
}

// 用户角色
type UserRoleDao interface {
	FindRIdsByUId(ctx context.Context, uId int64) ([]int64, error)
}

// 角色资源
type RoleResourceDao interface {
	FindResourceByRoles(ctx context.Context, rIds []int64) ([]int64, error)
}
