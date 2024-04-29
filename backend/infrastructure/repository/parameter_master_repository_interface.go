package repository

import (
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
)

type ParameterMasterRepositoryInterface interface {
	// 全パラメータマスタの取得
	ReadParameterMasters(query interface{}, args ...interface{}) ([]entity.ParameterMaster, *exception.CustomException)
}
