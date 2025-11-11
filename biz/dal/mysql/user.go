package mysql

import (
	"CarBuyerAssitance/pkg/constants"
	"CarBuyerAssitance/pkg/errno"
	"context"
	"errors"
	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, user *User) (string, error) {
	err := db.WithContext(ctx).
		Table(constants.TableUser).
		Create(user).
		Error
	if err != nil {
		return "", errno.NewErrNo(errno.InternalDatabaseErrorCode, "Create User Error:"+err.Error())
	}
	return user.UserId, nil
}

func IsUserExist(ctx context.Context, user_id string) (bool, error) {
	var userInfo *User
	err := db.WithContext(ctx).
		Table(constants.TableUser).
		Where("user_id = ?", user_id).
		First(&userInfo).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //没找到了说明用户不存在
			return false, nil
		}
		return false, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query user: %v", err)
	}
	return true, nil
}

func GetUserInfoByRoleId(ctx context.Context, role_id string) (*User, error) {
	var userInfo *User
	err := db.WithContext(ctx).
		Table(constants.TableUser).
		Where("user_id = ?", role_id).
		First(&userInfo).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "query user Info error:"+err.Error())
	}
	return userInfo, nil
}
