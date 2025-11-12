package pack

import (
	"CarBuyerAssitance/biz/dal/mysql"
	resp "CarBuyerAssitance/biz/model/model"
	"CarBuyerAssitance/biz/service/model"
	"math/rand"
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
		Status:         int64(user.Status),
		Address:        user.Address,
		CreatedAt:      user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      user.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:      strconv.FormatInt(0, 10),
	}
}

func ConsultResult(data *model.ConsultResult, id int) *resp.ConsultResult {
	result := &resp.ConsultResult{
		Analysis:  data.Analysis,
		Proposal:  data.Proposal,
		ConsultID: int64(id),
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
func Consultation(data *model.AllConsulation) *resp.Consultation {
	if data == nil {
		return nil
	}

	result := &resp.Consultation{
		Consult: &resp.Consult{
			UserId:          data.Consultation.UserId,
			ConsultId:       int64(data.Consultation.ConsultId),
			BudgetRange:     data.Consultation.BudgetRange,
			PreferredType:   data.Consultation.PreferredType,
			UseCase:         data.Consultation.UseCase,
			FuelType:        data.Consultation.FuelType,
			BrandPreference: data.Consultation.BrandPreference,
		},
	}

	// 转换ConsultResult
	if data.ConsultResult.Result != nil {
		result.ConsultResult = &resp.ConsultResult{
			Analysis: data.ConsultResult.Analysis,
			Proposal: data.ConsultResult.Proposal,
			Result:   make([]*resp.Car, len(data.ConsultResult.Result)),
		}

		// 转换Car列表
		for i, car := range data.ConsultResult.Result {
			result.ConsultResult.Result[i] = &resp.Car{
				ImageUrl:          car.ImageUrl,
				CarName:           car.CarName,
				FuelConsumption:   car.FuelConsumption,
				Power:             car.Power,
				Seat:              car.Seat,
				Drive:             car.Drive,
				RecommendedReason: car.RecommendedReason,
			}
		}
	} else {
		result.ConsultResult = &resp.ConsultResult{
			Analysis: data.ConsultResult.Analysis,
			Proposal: data.ConsultResult.Proposal,
			Result:   []*resp.Car{},
		}
	}

	return result
}

func PointList(data []*mysql.Points) *resp.PointList {
	if data == nil {
		return &resp.PointList{
			Item: []*resp.Point{},
			Num:  0,
			Sum:  0,
		}
	}

	// 计算总积分
	var totalSum int64 = 0
	for _, point := range data {
		totalSum += int64(point.Points)
	}

	// 转换每个Point
	points := make([]*resp.Point, len(data))
	for i, point := range data {
		points[i] = &resp.Point{
			PointID:   int64(point.PointID),
			UserID:    point.UserID,
			Points:    int64(point.Points),
			Reason:    point.Reason,
			CreatedAt: point.CreateTime.Format("2006-01-02 15:04:05"),
			UpdatedAt: point.UpdateTime.Format("2006-01-02 15:04:05"),
		}
	}

	return &resp.PointList{
		Item: points,
		Num:  int64(len(data)),
		Sum:  totalSum,
	}
}

func Gift(data []*mysql.Gift) *resp.GiftList {
	if data == nil {
		return &resp.GiftList{
			Item:  []*resp.Gift{},
			Total: 0,
		}
	}

	// 转换每个Gift
	gifts := make([]*resp.Gift, len(data))
	for i, gift := range data {
		// 将bool转换为int64
		var isOnline int64 = 0
		if gift.IsOnline {
			isOnline = 1
		}

		gifts[i] = &resp.Gift{
			GiftID:         gift.GiftID,
			GiftName:       gift.GiftName,
			RequiredPoints: int64(gift.RequiredPoints),
			StockQuantity:  int64(gift.StockQuantity),
			CoverImageURL:  gift.CoverImageURL,
			IsOnline:       isOnline,
			CreatedAt:      gift.CreateTime.Format("2006-01-02 15:04:05"),
			UpdatedAt:      gift.UpdateTime.Format("2006-01-02 15:04:05"),
		}
	}

	return &resp.GiftList{
		Item:  gifts,
		Total: int64(len(data)),
	}
}

func Order(data *mysql.Exchange) *resp.Order {
	return &resp.Order{
		UserID:     data.UserId,
		GiftName:   data.GiftName,
		NeedPoints: int64(data.NeedPoints),
		Status:     int64(data.Status),
		OrderTime:  data.ExchangeTime.Format("2006-01-02 15:04:05"),
		Id:         data.ExchangeId,
		Name:       data.Name,
		Address:    data.Address,
		Phone:      data.Phone,
	}
}

func CList(data []*model.AllConsulation, total int64) *resp.ConsultationList {
	result := make([]*resp.Consultation, 0)
	for _, v := range data {
		result = append(result, Consultation(v))
	}
	return &resp.ConsultationList{
		Total: total,
		Item:  result,
	}
}

func Budget() *resp.BudgetList {
	r1 := &resp.Budget{
		BudgetName: "5w——10w",
		Value:      int64(rand.Uint32()%3 + 14), // 14~16
	}
	r2 := &resp.Budget{
		BudgetName: "10w——20w",
		Value:      int64(rand.Uint32()%11 + 10), // 10~20
	}
	r3 := &resp.Budget{
		BudgetName: "20w以上",
		Value:      int64(rand.Uint32()%4 + 1), // 20~30
	}

	result := make([]*resp.Budget, 0)
	result = append(result, r1, r2, r3)

	return &resp.BudgetList{
		Item:  result,
		Total: 3,
	}
}

func Frequency() *resp.FrequencyList {
	f1 := &resp.Frequency{
		FrequencyName: "0-5次",
		Value:         int64(rand.Uint32()%6 + 15), // 15~20
	}
	f2 := &resp.Frequency{
		FrequencyName: "5-20次",
		Value:         int64(rand.Uint32()%11 + 10), // 10~20
	}
	f3 := &resp.Frequency{
		FrequencyName: "20次以上",
		Value:         int64(rand.Uint32()%6 + 5), // 5~10
	}

	result := make([]*resp.Frequency, 0)
	result = append(result, f1, f2, f3)

	return &resp.FrequencyList{
		Item:  result,
		Total: 3,
	}
}

func Scene() *resp.SceneList {
	s1 := &resp.Scene{
		SceneName: "通勤",
		Value:     int64(rand.Uint32()%11 + 20), // 20~30
	}
	s2 := &resp.Scene{
		SceneName: "家庭",
		Value:     int64(rand.Uint32()%11 + 15), // 15~25
	}
	s3 := &resp.Scene{
		SceneName: "商务",
		Value:     int64(rand.Uint32()%11 + 10), // 10~20
	}
	s4 := &resp.Scene{
		SceneName: "其他",
		Value:     int64(rand.Uint32()%6 + 5), // 5~10
	}

	result := make([]*resp.Scene, 0)
	result = append(result, s1, s2, s3, s4)

	return &resp.SceneList{
		Item:  result,
		Total: 4,
	}
}

func OrderList(data []*mysql.Exchange, total int64) *resp.OrderList {
	res := make([]*resp.Order, 0)
	for _, v := range data {
		res = append(res, Order(v))
	}
	return &resp.OrderList{
		Item:  res,
		Total: total,
	}
}
