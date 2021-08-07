package define

// 用户注册
type UserRegisterReq struct {
	UserServiceSignUp
	Passport string `v:"required#请输入账号"` // 账号
	Password string `v:"required#请输入密码"` // 密码
	Nickname string `v:"required#请输入昵称"` // 昵称
}

// 注册输入参数
type UserServiceSignUp struct {
	Passport string
	Password string
	Nickname string
}

// 用户信息
type UserGetProfileOutput struct {
	Id       uint
	Nickname string
	Passport string
}

// 修改用户
type UserUpdateProfileReq struct {
	UserUpdateProfileInput
	Nickname string `v:"required#请输入昵称信息"` // 昵称
	Passport string `v:"required#请输入账号"`   // 账号
}

// 修改用户信息
type UserUpdateProfileInput struct {
	Id       uint
	Nickname string
	Passport string
}

// 查询用户信息
type UserProfileReq struct {
	Id uint `v:"min:1#id是必须的"`
}

type UserServiceProfile struct {
	Id uint
}
