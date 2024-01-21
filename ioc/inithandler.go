package ioc

import (
	"XiaoBaoSecurity/com/authority"
	"XiaoBaoSecurity/repository"
	"XiaoBaoSecurity/repository/dao"
	"XiaoBaoSecurity/service"
	"XiaoBaoSecurity/utils"
	"XiaoBaoSecurity/web"
)

// InitUserHandle 构造老复杂了
func InitUserHandle(jwtUtils utils.JWTUtils) web.UserHandler {
	db := InitDB()
	userDao := dao.NewUserDao(db)
	userRoleDao := dao.NewUserRoleDao(db)
	roleResourceDao := dao.NewRoleResourceDao(db)
	resourceDao := dao.NewResourceDao(db)
	userRepository := repository.
		NewDefaultUserRepository(userDao, userRoleDao, roleResourceDao, resourceDao)
	userService := service.NewDefaultUserService(userRepository)
	authorityHandler := authority.NewAuthority(jwtUtils)
	handler := web.NewUserHandler(userService, authorityHandler)
	return handler
}
