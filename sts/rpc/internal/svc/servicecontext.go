package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"test/sts/rpc/internal/config"
	"test/sts/rpc/model"
)

type ServiceContext struct {
	Config   config.Config
	DataBase model.StsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		DataBase: model.NewStsModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
