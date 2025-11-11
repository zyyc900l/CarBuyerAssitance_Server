package service

import (
	"CarBuyerAssitance/biz/service/model"
	"CarBuyerAssitance/pkg/utils"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
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
	consultResult, err := utils.CallOpenAIWithConsult(svc.ctx, consult)
	if err != nil {
		return nil, fmt.Errorf("call openai error:" + err.Error())
	}
	return consultResult, nil
}
