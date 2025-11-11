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
func DeleteUser(ctx context.Context, user_id string) error {
	err := db.WithContext(ctx).
		Table(constants.TableUser).
		Where("user_id = ?", user_id).
		Update("status", -1).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "delete user Info error:"+err.Error())
	}
	return nil
}

func UpdateUserInfo(ctx context.Context, user *User) (*User, error) {
	updateData := map[string]interface{}{
		"budget_min":      user.BudgetMin,
		"budget_max":      user.BudgetMax,
		"preferred_type":  user.PreferredType,
		"preferred_brand": user.PreferredBrand,
		"address":         user.Address,
	}

	err := db.WithContext(ctx).
		Table(constants.TableUser).
		Where("user_id = ?", user.UserId).
		Updates(updateData).Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "update user info error: "+err.Error())
	}

	// 查询更新后的用户数据
	var updatedUser *User
	err = db.WithContext(ctx).
		Table(constants.TableUser).
		Where("user_id = ?", user.UserId).
		Find(&updatedUser).Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "update user info error: "+err.Error())
	}
	return updatedUser, nil

}
