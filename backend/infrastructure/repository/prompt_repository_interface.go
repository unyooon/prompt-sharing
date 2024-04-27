package repository

import (
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
)

type PromptRepositoryInterface interface {
	Create(prompt entity.Prompt) (entity.Prompt, *exception.CustomException)
	ReadPrompts(limit int, offset int, query interface{}, args ...interface{}) ([]entity.Prompt, *exception.CustomException)
}
