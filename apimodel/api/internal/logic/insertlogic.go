package logic

import (
	"context"
	"test/sts/rpc/types/sts"

	"test/apimodel/api/internal/svc"
	"test/apimodel/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertLogic) Insert(req *types.InsertReq) (resp *types.InsertReply, err error) {
	// todo: add your logic here and delete this line
	con, err := l.svcCtx.StsRpc.Insert(l.ctx, &sts.InsertRequest{
		Title: req.Title,
		Text:  req.Text,
	})
	if err != nil {
		return nil, err
	}
	return &types.InsertReply{
		Id: con.Id,
	}, nil
}
