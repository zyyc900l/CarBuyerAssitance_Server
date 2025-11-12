package service

import (
	"CarBuyerAssitance/biz/dal/mysql"
	"CarBuyerAssitance/pkg/crypt"
	"CarBuyerAssitance/pkg/errno"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{
		ctx: ctx,
		c:   c,
	}
}

func (svc *UserService) Register(user *mysql.User) (string, error) {
	var err error
	exist, err := mysql.IsUserExist(svc.ctx, user.UserId)
	if err != nil {
		return "", err
	}
	if exist {
		return "", errno.NewErrNo(errno.ServiceUserExistCode, "user  exist")
	}
	user.Password, err = crypt.PasswordHash(user.Password)
	if err != nil {
		return "", err
	}
	return mysql.CreateUser(svc.ctx, user)
}
func (svc *UserService) Login(user_id string, password string) (*mysql.User, error) {
	exist, err := mysql.IsUserExist(svc.ctx, user_id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceUserNotExistCode, "user not exist")
	}
	userInfo, err := mysql.GetUserInfoByRoleId(svc.ctx, user_id)
	if err != nil {
		return nil, err
	}
	// 激活检验
	if userInfo.Status == -1 {
		return nil, errno.NewErrNo(errno.ServiceUserBanExistCode, "user have been deleted")
	}
	if !crypt.VerifyPassword(password, userInfo.Password) {
		return nil, errno.Errorf(errno.ServiceUserPasswordError, "password not match")
	}
	return userInfo, nil
}

func (svc *UserService) Query(user_id string) (*mysql.User, error) {
	exist, err := mysql.IsUserExist(svc.ctx, user_id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceUserNotExistCode, "user not exist")
	}
	userInfo, err := mysql.GetUserInfoByRoleId(svc.ctx, user_id)
	if err != nil {
		return nil, err
	}
	// 激活检验
	if userInfo.Status == -1 {
		return nil, errno.NewErrNo(errno.ServiceUserBanExistCode, "user have been deleted")
	}
	return userInfo, nil
}
func (svc *UserService) Update(user *mysql.User) (*mysql.User, error) {
	user.UserId = GetUserIDFromContext(svc.c)
	exist, err := mysql.IsUserExist(svc.ctx, user.UserId)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceUserNotExistCode, "user not exist")
	}
	userInfo, err := mysql.GetUserInfoByRoleId(svc.ctx, user.UserId)
	if err != nil {
		return nil, err
	}
	// 激活检验
	if userInfo.Status == -1 {
		return nil, errno.NewErrNo(errno.ServiceUserBanExistCode, "user have been deleted")
	}
	I, err := mysql.UpdateUserInfo(svc.ctx, user)
	if err != nil {
		return nil, err
	}
	return I, nil
}
