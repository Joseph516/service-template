package service

import (
	"context"
	"service-template/global"
	"service-template/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
		dao: dao.NewDao(global.DBEngine),
	}
	return svc
}
