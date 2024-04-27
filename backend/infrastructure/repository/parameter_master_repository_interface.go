package repository

import (
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
)

type ParameterMasterRepositoryInterface interface {
	// 全パラメータマスタの取得
	ReadParameterMasters() ([]entity.ParameterMaster, *exception.CustomException)
}
