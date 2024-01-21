package web

import (
	"XiaoBaoSecurity/com/authority"
	"XiaoBaoSecurity/domian"
	"XiaoBaoSecurity/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultUserHandler struct {
	us service.UserService
	aa authority.AuthorityHandler
}

func NewUserHandler(us service.UserService,
	aa authority.AuthorityHandler) UserHandler {
	return &DefaultUserHandler{
		us: us,
		aa: aa,
	}
}
func (d *DefaultUserHandler) Register(server *gin.Engine) {
	s := server.Group("/users")
	s.POST("/signup", d.SignUp)
	s.POST("/login", d.aa.SetAuthority(d.Login))
	s.GET("/ping", d.aa.CheckStaticAuthority(d.Ping))
	s.GET("/ping/:id", d.aa.CheckDynamicsAuthority(d.PingId))
}
func (d *DefaultUserHandler) SignUp(ctx *gin.Context) {
	//注册
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	var req SignUpReq
	// Bind 方法会根据 Content-Type 来解析你的数据到 req 里面
	// 解析错了，就会直接写回一个 400 的错误
	if err := ctx.Bind(&req); err != nil {
		return
	}
	//暂时不做邮箱的校验了
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}
	err := d.us.SignUp(ctx, req.Email, req.Password)
	if err != nil {
		ctx.String(200, "注册失败")
		return
	}
	ctx.String(http.StatusOK, "注册成功")
}

func (d *DefaultUserHandler) Login(ctx *gin.Context) (domian.AuthorityUserInfo, error) {
	//注册
	type SignUpReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return domian.AuthorityUserInfo{}, errors.New("参数解析失败")
	}
	return d.us.LoginByEmail(ctx, req.Email, req.Password)
}

func (d *DefaultUserHandler) Ping(ctx *gin.Context) {
	ctx.String(200, "pang")
}
func (d *DefaultUserHandler) PingId(ctx *gin.Context) {
	ctx.String(200, "pang-带资源的路由")
}
