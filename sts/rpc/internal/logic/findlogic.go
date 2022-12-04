package logic

import (
	"context"

	"test/sts/rpc/internal/svc"
	"test/sts/rpc/types/sts"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLogic {
	return &FindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindLogic) Find(in *sts.FindRequest) (*sts.FindResponse, error) {
	// todo: add your logic here and delete this line
	conn, err := l.svcCtx.DataBase.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &sts.FindResponse{Text: conn.Text, Title: conn.Title}, nil
}
