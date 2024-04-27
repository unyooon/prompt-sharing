package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/domain/dto"
	"github.com/unyooon/prompt-sharing/domain/exception"
)

type PromptsReader interface {
	Execute(c *gin.Context) (dto.ReadPromptsResponse, *exception.CustomException)
}
