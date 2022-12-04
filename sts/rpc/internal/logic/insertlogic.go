package logic

import (
	"context"
	"test/sts/rpc/model"

	"test/sts/rpc/internal/svc"
	"test/sts/rpc/types/sts"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertLogic) Insert(in *sts.InsertRequest) (*sts.InsertResponse, error) {
	// todo: add your logic here and delete this line
	conn, err := l.svcCtx.DataBase.Insert(l.ctx, &model.Sts{
		Title: in.Title,
		Text:  in.Text,
	})
	if err != nil {
		return nil, err
	}
	id, err := conn.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &sts.InsertResponse{Id: id}, nil
}
