package service

import (
	"CarBuyerAssitance/biz/dal/mysql"
	"CarBuyerAssitance/biz/service/model"
	"CarBuyerAssitance/pkg/crypt"
	"CarBuyerAssitance/pkg/errno"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type AdminService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAdminService(ctx context.Context, c *app.RequestContext) *AdminService {
	return &AdminService{
		ctx: ctx,
		c:   c,
	}
}
func (svc *AdminService) AddUser(user *mysql.User) (string, error) {
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
func (svc *AdminService) DeleteUser(user_id string) error {
	var err error
	exist, err := mysql.IsUserExist(svc.ctx, user_id)
	hlog.Info(exist)
	if err != nil {
		return err
	}
	if !exist {
		return errno.NewErrNo(errno.ServiceUserExistCode, "user not exist")
	}
	err = mysql.DeleteUser(svc.ctx, user_id)
	return err
}

func (svc *AdminService) QueryAllConsult(page int, page_size int) ([]*model.AllConsulation, int64, error) {
	return mysql.QueryAllConsultMessages(svc.ctx, page, page_size)
}
