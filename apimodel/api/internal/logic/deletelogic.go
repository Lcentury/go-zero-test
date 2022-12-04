package logic

import (
	"context"
	"test/sts/rpc/types/sts"

	"test/apimodel/api/internal/svc"
	"test/apimodel/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteReq) (resp *types.DeleteReply, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.StsRpc.Delete(l.ctx, &sts.DeleteRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.DeleteReply{
		Rep: "yes",
	}, nil
}
