package repository

import (
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
)

type LlmMasterRepositoryInterface interface {
	ReadLlmMaster(query interface{}, args ...interface{}) (entity.LlmMaster, *exception.CustomException)
	ReadLlmsMaster(query interface{}, args ...interface{}) ([]entity.LlmMaster, *exception.CustomException)
}
