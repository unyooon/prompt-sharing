package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/usecase"
)

type PromptController struct {
	PromptCreater usecase.PromptCreater
	PromptReader  usecase.PromptsReader
}

func NewPromptController(
	promptCreater usecase.PromptCreater,
	promptReader usecase.PromptsReader,
) *PromptController {
	return &PromptController{
		PromptCreater: promptCreater,
		PromptReader:  promptReader,
	}
}

// ReadPrompts godoc
// @Summary      Read Prompts
// @Description  read prompts
// @Tags         prompt
// @Accept       json
// @Produce      json
// @Param        request query dto.ReadPromptsRequest true "request"
// @Success      200 {object} dto.ReadPromptsResponse
// @Router       /prompts [get]
func (pc *PromptController) ReadPrompts(c *gin.Context) {
	res, err := pc.PromptReader.Execute(c)
	Handler(c, res, err)
}

// CreatePrompt godoc
// @Summary      Create Prompt
// @Description  create prompt
// @Tags         prompt
// @Accept       json
// @Produce      json
// @Param        prompt body dto.CreatePromptRequest true "prompt"
// @Success      200
// @Router       /prompts [post]
func (pc *PromptController) CreatePrompt(c *gin.Context) {
	err := pc.PromptCreater.Execute(c)
	Handler(c, nil, err)
}
