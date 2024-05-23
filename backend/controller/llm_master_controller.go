package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/usecase"
)

type LlmMasterController struct {
	LlmMasterReader usecase.LlmMastersReader
}

func NewLlmMasterController(
	llmMasterReader usecase.LlmMastersReader,
) *LlmMasterController {
	return &LlmMasterController{
		LlmMasterReader: llmMasterReader,
	}
}

// ReadLlmMasters godoc
// @Summary      Read LLM Masters
// @Description  Read LLM Masters
// @Tags         llm_master
// @Accept       json
// @Produce      json
// @Success      200 {array} dto.ReadLlmMasterResponse
// @Router       /llm-masters [get]
func (lmc *LlmMasterController) ReadLlmMasters(c *gin.Context) {
	res, err := lmc.LlmMasterReader.Execute(c)
	Handler(c, res, err)
}
