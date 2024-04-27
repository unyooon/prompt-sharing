package dto

type ReadPromptsResponse struct {
	Page BasePageResponse
	Data []ReadPromptResponse
}

type ReadPromptResponse struct {
	ID          uint   `json:"id"`
	LlmName     string `json:"llmName"`
	Parameters  []ParameterResponse
	Text        string `json:"text"`
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
	CreatedUser CreatedUserResponse
	CreatedAt   string `json:"createdAt"`
}

type CreatedUserResponse struct {
	DisplayName string `json:"displayName"`
}

type ParameterResponse struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
