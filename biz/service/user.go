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
	if !exist {
		return "", errno.NewErrNo(errno.ServiceUserNotExistCode, "user not exist")
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
	if !crypt.VerifyPassword(password, userInfo.Password) {
		return nil, errno.Errorf(errno.ServiceUserPasswordError, "password not match")
	}
	return userInfo, nil
}
