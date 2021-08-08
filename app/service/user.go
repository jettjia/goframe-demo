package service

import (
	"context"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"my-app/app/dao"
	"my-app/app/define"
	"my-app/app/model"
)

var User = userService{}

type userService struct {}

// 用户注册
func (s *userService) SignUp(ctx context.Context, input define.UserServiceSignUp) error  {
	return dao.SysUser.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var user *model.SysUser
		if err := gconv.Struct(input, &user); err != nil {
			return err
		}
		// 检测给定的账号是否唯一
		if err := s.CheckPassportUnique(ctx, input.Passport); err != nil {
			return err
		}
		// 检测给定的昵称是否唯一
		if err := s.CheckNicknameUnique(ctx, input.Nickname); err != nil {
			return err
		}
		// 添加用户
		user.Password = s.EncryptPassword(input.Password, input.Password)
		_, err := dao.SysUser.Ctx(ctx).Data(user).OmitEmpty().Save()
		return err
	})
}

// 获取个人信息
func (s *userService) GetProfileById(ctx context.Context, id int) (output *define.UserGetProfileOutput, err error) {
	output = &define.UserGetProfileOutput{}
	if err := dao.SysUser.Ctx(ctx).WherePri(id).Scan(output); err != nil {
		return nil, err
	}
	return
}

// 修改个人资料
func (s *userService) UpdateProfile(ctx context.Context, input define.UserUpdateProfileInput) error {
	return dao.SysUser.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var (
			err error
		)
		n, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.C.Nickname, input.Nickname).
			WhereNot(dao.SysUser.C.Id, input.Id).Count()
		if err != nil {
			return err
		}
		if n > 0 {
			return gerror.Newf(`昵称"%s"已被占用`, input.Nickname)
		}

		_, err = dao.SysUser.Ctx(ctx).OmitEmpty().Data(input).Where(dao.SysUser.C.Id, input.Id).Update()
		return err
	})
}

// Login 用户登录逻辑
func (*userService)Login(ctx context.Context, username string) (output *define.UserGetProfileOutput, err error) {
	output = &define.UserGetProfileOutput{}
	if err := dao.SysUser.Ctx(ctx).Where("nickname=?", username).Scan(output); err != nil {
		return nil, err
	}
	return
}

// 检测给定的账号是否唯一
func (s *userService) CheckPassportUnique(ctx context.Context, passport string) error {
	n, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.C.Passport, passport).Count()
	if err != nil {
		return err
	}
	if n > 0  {
		return gerror.Newf(`账号"%s"已被占用`, passport)
	}
	return nil
}

// 检测给定的昵称是否唯一
func (s *userService) CheckNicknameUnique(ctx context.Context, nickname string) error {
	n, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.C.Nickname, nickname).Count()
	if err != nil {
		return err
	}

	if n > 0 {
		return gerror.Newf("昵称%已经被占用", nickname)
	}
	return nil
}

// 将密码按照内部算法进行加密
func (s *userService) EncryptPassword(passport, password string) string {
	return gmd5.MustEncrypt(passport + password)
}