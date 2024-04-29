package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/unyooon/prompt-sharing/core/constants"
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
	"github.com/unyooon/prompt-sharing/infrastructure/db"
)

type LlmMasterRepository struct {
	db *gorm.DB
}

func NewLlmMasterRepository(db db.DbInterface) *LlmMasterRepository {
	d := db.Connect()
	return &LlmMasterRepository{
		db: d,
	}
}

// LLMマスタの単体取得
func (repo *LlmMasterRepository) ReadLlmMaster(query interface{}, args ...interface{}) (entity.LlmMaster, *exception.CustomException) {
	var llmMaster entity.LlmMaster
	if err := repo.db.Where(query, args).First(&llmMaster).Error; err != nil {
		return llmMaster, &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}
	return llmMaster, nil
}

// 全LLMマスタ取得
func (repo *LlmMasterRepository) ReadLlmsMaster(query interface{}, args ...interface{}) ([]entity.LlmMaster, *exception.CustomException) {
	var llmsMaster []entity.LlmMaster
	if err := repo.db.Where(query, args).Find(&llmsMaster).Error; err != nil {
		return llmsMaster, &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}
	return llmsMaster, nil
}
