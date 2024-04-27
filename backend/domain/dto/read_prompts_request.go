package dto

import (
	"github.com/unyooon/prompt-sharing/core/types"
)

type ReadPromptsRequest struct {
	types.Query
	LlmId    *uint `json:"llmId"`
	IsPublic bool  `json:"isPublic"`
}
