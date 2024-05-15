package dto

type CreatePromptRequest struct {
	LlmId       uint        `json:"llmId" binding:"required"`
	Parameters  []Parameter `json:"parameters" binding:"required"`
	Text        string      `json:"text" binding:"required"`
	Description string      `json:"description" binding:"required"`
}

type Parameter struct {
	ParameterMasterId uint `json:"parameterId"`
	Value             int  `json:"value"`
}
