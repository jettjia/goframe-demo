package api

import (
	"github.com/gogf/gf/net/ghttp"
	"my-app/app/define"
	"my-app/app/service"
	"my-app/library/response"
)

// 注册控制器
var User = userApi{}

type userApi struct {}

// @summary 执行注册提交处理
// @description 注意提交的密码是明文。
// @description 注册成功后返回结果
// @tags    后台台-注册
// @produce json
// @param   entity body define.UserRegisterReq true "请求参数" required
// @router  /user/registry [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) Register(r *ghttp.Request) {
	var (
		req *define.UserRegisterReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.User.SignUp(r.Context(), req.UserServiceSignUp); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 获取用户详情信息
// @tags    用户服务
// @produce json
// @router  /user/profile [GET]
// @success 200 {object} model.SysUser "用户信息"
func (a *userApi) Profile(r *ghttp.Request) {
	userId := r.GetInt("userId")
	info, err := service.User.GetProfileById(r.Context(), userId)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", info)
}

// @summary 修改用户详情信息
// @tags    用户服务
// @produce json
// @router  /user/update [post]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) UpdateProfile(r *ghttp.Request) {
	var (
		req *define.UserUpdateProfileReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.User.UpdateProfile(r.Context(), req.UserUpdateProfileInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}
