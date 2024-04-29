package dto

import (
	"github.com/unyooon/prompt-sharing/core/types"
)

type ReadPromptsRequest struct {
	types.Query
	LlmId    *uint `form:"llmId" binding:"required"`
	IsPublic bool  `form:"isPublic" binding:"required"`
}
