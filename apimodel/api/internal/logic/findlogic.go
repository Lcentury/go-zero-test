package logic

import (
	"context"
	"test/sts/rpc/types/sts"

	"test/apimodel/api/internal/svc"
	"test/apimodel/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLogic {
	return &FindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindLogic) Find(req *types.FindReq) (resp *types.FindReply, err error) {
	// todo: add your logic here and delete this line
	con, err := l.svcCtx.StsRpc.Find(l.ctx, &sts.FindRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.FindReply{
		Title: con.Title,
		Text:  con.Text,
	}, nil
}
