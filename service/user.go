package service

import (
	"XiaoBaoSecurity/domian"
	"XiaoBaoSecurity/repository"
	"XiaoBaoSecurity/repository/dao"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidUserOrPassword = errors.New("账号或密码错误")
)

type DefaultUserService struct {
	ur repository.UserRepository
}

func NewDefaultUserService(ur repository.UserRepository) UserService {
	return &DefaultUserService{
		ur: ur,
	}
}
func (d *DefaultUserService) SignUp(ctx context.Context, email string, password string) error {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return d.ur.Create(ctx, email, string(fromPassword))
}
func (d *DefaultUserService) LoginByEmail(ctx context.Context, email string, password string) (domian.AuthorityUserInfo, error) {
	user, err := d.ur.FindByEmail(ctx, email)
	if err != nil {
		return domian.AuthorityUserInfo{}, err
	}
	//去解密对比密码
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err = bcrypt.CompareHashAndPassword(encrypted, []byte(user.Password))
	if err != nil {
		//没通过呗,debug或者info日志
		return domian.AuthorityUserInfo{}, ErrInvalidUserOrPassword
	}
	//通过比对，登录成功，去查权限信息
	reslist, err := d.ur.FindUserAuthority(ctx, user.Id)
	if err != nil {
		return domian.AuthorityUserInfo{}, err
	}
	//都查到了，并且处理到统一格式了，没什么事情
	return d.buildAuthorityUserInfo(user, reslist), nil
}
func (d *DefaultUserService) buildAuthorityUserInfo(user dao.User, reslist []dao.Resource) domian.AuthorityUserInfo {
	var info domian.AuthorityUserInfo
	info.UId = user.Id
	uMap := make(map[string]byte)
	uList := make([]domian.UrlNode, 0)
	for i := 0; i < len(reslist); i++ {
		//是动态路由
		if reslist[i].IsDynamicRouting == 1 {
			//分类到动态路由的权限里面
			uList = append(uList, domian.UrlNode{
				Url:   reslist[i].DynamicUrl,
				Level: reslist[i].Level,
			})
		} else {
			uMap[reslist[i].Url] = 1 //基础路由
		}
	}
	info.UrlMap = uMap
	info.UrlList = uList
	return info
}
