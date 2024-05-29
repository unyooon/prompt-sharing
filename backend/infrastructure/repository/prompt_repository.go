package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/unyooon/prompt-sharing/core/constants"
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
	"github.com/unyooon/prompt-sharing/infrastructure/db"
)

type PromptRepository struct {
	db *gorm.DB
}

func NewPromptRepository(db db.DbInterface) *PromptRepository {
	d := db.Connect()
	return &PromptRepository{
		db: d,
	}
}

func (repo *PromptRepository) ReadPrompts(limit int, offset int, query interface{}, args ...interface{}) ([]entity.Prompt, *exception.CustomException) {
	var prompts []entity.Prompt

	if err := repo.db.Preload("Llm").Preload("Parameters").Where(query, args...).Order("created_at").Limit(limit).Offset(offset).Find(&prompts).Error; err != nil {
		return prompts, &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}
	return prompts, nil
}

func (repo *PromptRepository) Create(prompt entity.Prompt) (entity.Prompt, *exception.CustomException) {
	if err := repo.db.Create(&prompt).Error; err != nil {
		return prompt, &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}
	return prompt, nil
}

func (repo *PromptRepository) Delete(prompt entity.Prompt) *exception.CustomException {
	// トランザクションの開始
	tx := repo.db.Begin()

	// トランザクションのロールバック
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// Promptに紐づいているParametersを取得
	var parameters []entity.Parameter
	if err := tx.Model(&prompt).Association("Parameters").Find(&parameters); err != nil {
		tx.Rollback()
		return &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err.Error,
		}
	}

	// Promptに紐づいているParametersを削除
	if err := tx.Unscoped().Delete(&parameters).Error; err != nil {
		tx.Rollback()
		return &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}

	// Promptを削除
	if err := tx.Unscoped().Delete(&prompt).Error; err != nil {
		tx.Rollback()
		return &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}

	// トランザクションのコミット
	if err := tx.Commit().Error; err != nil {
		return &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}

	return nil
}
