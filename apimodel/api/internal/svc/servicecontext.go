package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"test/apimodel/api/internal/config"
	"test/sts/rpc/stsclient"
)

type ServiceContext struct {
	Config config.Config
	StsRpc stsclient.Sts
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		StsRpc: stsclient.NewSts(zrpc.MustNewClient(c.StsRpc)),
	}
}
