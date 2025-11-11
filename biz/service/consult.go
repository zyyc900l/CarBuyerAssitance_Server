package service

import (
	"CarBuyerAssitance/biz/dal/mysql"
	"CarBuyerAssitance/biz/service/model"
	"CarBuyerAssitance/biz/service/taskqueue"
	"CarBuyerAssitance/pkg/constants"
	"CarBuyerAssitance/pkg/errno"
	"CarBuyerAssitance/pkg/utils"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"math/rand"
	"strings"
	"time"
)

type ConsultService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewConsultService(ctx context.Context, c *app.RequestContext) *ConsultService {
	return &ConsultService{
		ctx: ctx,
		c:   c,
	}
}

func (svc *ConsultService) Consult(consult *model.Consult) (*model.ConsultResult, error) {
	id := GetUserIDFromContext(svc.c)
	con, err := mysql.CreateConsultation(svc.ctx, id, consult)
	if err != nil {
		return nil, err
	}
	consultResult, err := utils.CallOpenAIWithConsult(svc.ctx, consult)
	if err != nil {
		return nil, fmt.Errorf("call openai error:" + err.Error())
	}
	for i, v := range consultResult.Result {
		if strings.HasPrefix(v.ImageUrl, constants.UrlPrefix) { //当ai没找到图片
			consultResult.Result[i].ImageUrl = defaultImage()
		}
	}
	err = mysql.SaveConsultResult(svc.ctx, con.ConsultId, consultResult)
	if err != nil {
		return nil, err
	}
	p := &mysql.Points{
		UserID: id,
		Points: 5,
		Reason: "咨询",
	}
	taskqueue.AddUpdateScoreTask(svc.ctx, constants.TaskQueue, p)
	return consultResult, nil
}

func (svc *ConsultService) QueryConsult(consult_id int) (*model.AllConsulation, error) {
	return mysql.QueryConsultMessage(svc.ctx, consult_id)
}

func (svc *ConsultService) QueryUserPoint() ([]*mysql.Points, error) {
	id := GetUserIDFromContext(svc.c)
	return mysql.GetUserPoints(svc.ctx, id)
}

func (svc *ConsultService) QueryGift() ([]*mysql.Gift, error) {
	return mysql.GetOnlineGifts(svc.ctx)

}
func (svc *ConsultService) BuyGift(gift_id int64) (*mysql.Exchange, error) {
	uid := GetUserIDFromContext(svc.c)
	exist, err := mysql.IsGiftExist(svc.ctx, gift_id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceGiftNotExistCode, "gift not exist")
	}
	giftInfo, err := mysql.QueryGiftById(svc.ctx, gift_id)
	if err != nil {
		return nil, err
	}
	if giftInfo.StockQuantity <= 0 {
		return nil, errno.NewErrNo(errno.ServiceGiftRunOutExistCode, "gift no available")
	}
	pointsInfo, err := mysql.GetUserPoints(svc.ctx, uid)
	if err != nil {
		return nil, err
	}
	var sum int
	for _, p := range pointsInfo {
		sum += p.Points
	}
	if sum < giftInfo.RequiredPoints {
		return nil, errno.NewErrNo(errno.ServicePointRunOutExistCode, "point not enough")
	}
	err = mysql.BuyGift(svc.ctx, gift_id)
	if err != nil {
		return nil, err
	}
	i := &mysql.Exchange{
		GiftName:     giftInfo.GiftName,
		UserID:       uid,
		NeedPoints:   giftInfo.RequiredPoints,
		Status:       1,
		ExchangeTime: time.Now(),
	}
	info, err := mysql.CreateExchange(svc.ctx, i)
	p := &mysql.Points{
		UserID: uid,
		Points: -1 * giftInfo.RequiredPoints,
		Reason: "购买周边",
	}
	taskqueue.AddUpdateScoreTask(svc.ctx, constants.TaskQueue, p)
	return info, nil
}
func defaultImage() string {
	images := []string{
		constants.SmallCar,
		constants.MiddleCar,
		constants.MpvCar,
		constants.RunningCar,
	}

	// 随机选择一个图片
	randIndex := rand.Intn(len(images))
	return images[randIndex]
}
