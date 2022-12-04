package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StsModel = (*customStsModel)(nil)

type (
	// StsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStsModel.
	StsModel interface {
		stsModel
	}

	customStsModel struct {
		*defaultStsModel
	}
)

// NewStsModel returns a model for the database table.
func NewStsModel(conn sqlx.SqlConn) StsModel {
	return &customStsModel{
		defaultStsModel: newStsModel(conn),
	}
}
