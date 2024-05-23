package dto

import (
	"github.com/unyooon/prompt-sharing/core/types"
)

type ReadPromptsRequest struct {
	types.Query
	LlmId       *uint  `form:"llmId"`
	IsPublic    *bool  `form:"isPublic"`
	SearchQuery string `form:"q"`
}
