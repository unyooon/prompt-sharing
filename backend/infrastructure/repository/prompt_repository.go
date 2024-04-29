package repository

import (
	"github.com/jinzhu/gorm"
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
	// var prompts []entity.Prompt
	// if err := repo.db.Preload("Llm").Preload("ParameterMaster").Limit(limit).Offset(offset).Where(query, args).Error; err != nil {
	// 	return prompts, &exception.CustomException{
	// 		Code:    constants.InternalServerErrorCode,
	// 		Message: constants.DatabaseError,
	// 		Err:     err,
	// 	}
	// }
	// return prompts, nil

	// TODO: mock
	prompts := []entity.Prompt{
		{
			LlmId:       1,
			Text:        "prompt1",
			Description: "description1",
			IsPublic:    true,
			CreatedBy:   "user1",
		},
	}
	return prompts, nil
}

func (repo *PromptRepository) Create(prompt entity.Prompt) (entity.Prompt, *exception.CustomException) {
	// if err := repo.db.Create(&prompt).Error; err != nil {
	// 	return prompt, &exception.CustomException{
	// 		Code:    constants.InternalServerErrorCode,
	// 		Message: constants.DatabaseError,
	// 		Err:     err,
	// 	}
	// }
	// return prompt, nil

	// TODO: mock
	return prompt, nil
}
