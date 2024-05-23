package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/domain/dto"
	"github.com/unyooon/prompt-sharing/domain/exception"
)

type LlmMastersReader interface {
	Execute(c *gin.Context) ([]dto.ReadLlmMasterResponse, *exception.CustomException)
}
