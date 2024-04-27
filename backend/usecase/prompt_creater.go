package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/domain/exception"
)

type PromptCreater interface {
	Execute(c *gin.Context) *exception.CustomException
}
