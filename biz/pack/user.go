package pack

import (
	"CarBuyerAssitance/biz/dal/mysql"
	resp "CarBuyerAssitance/biz/model/model"
	"CarBuyerAssitance/biz/service/model"
	"strconv"
)

func User(user *mysql.User) *resp.UserInfo {
	return &resp.UserInfo{
		Username:       user.Username,
		UserId:         user.UserId,
		Phone:          user.Phone,
		BudgeMin:       user.BudgetMin,
		BudgeMax:       user.BudgetMax,
		PreferredType:  user.PreferredType,
		PreferredBrand: user.PreferredBrand,
		CreatedAt:      strconv.FormatInt(user.CreatedAt.Unix(), 10),
		UpdatedAt:      strconv.FormatInt(user.UpdatedAt.Unix(), 10),
		DeletedAt:      strconv.FormatInt(0, 10),
	}
}

func ConsultResult(data *model.ConsultResult) *resp.ConsultResult {
	result := &resp.ConsultResult{
		Analysis: data.Analysis,
		Proposal: data.Proposal,
	}

	// 转换Car列表
	if data.Result != nil {
		result.Result = make([]*resp.Car, len(data.Result))
		for i, car := range data.Result {
			result.Result[i] = &resp.Car{
				ImageUrl:          car.ImageUrl,
				CarName:           car.CarName,
				FuelConsumption:   car.FuelConsumption,
				Power:             car.Power,
				Seat:              car.Seat,
				Drive:             car.Drive,
				RecommendedReason: car.RecommendedReason,
			}
		}
	}

	return result
}
