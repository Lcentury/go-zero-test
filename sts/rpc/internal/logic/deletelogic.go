package logic

import (
	"context"

	"test/sts/rpc/internal/svc"
	"test/sts/rpc/types/sts"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *sts.DeleteRequest) (*sts.DeleteResponse, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.DataBase.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &sts.DeleteResponse{Rps: "Yes"}, nil
}
