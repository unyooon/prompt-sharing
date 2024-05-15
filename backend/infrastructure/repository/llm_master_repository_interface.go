package repository

import (
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
)

type LlmMasterRepositoryInterface interface {
	ReadLlmMaster(llmId uint) (entity.LlmMaster, *exception.CustomException)
	ReadLlmsMaster() ([]entity.LlmMaster, *exception.CustomException)
}
