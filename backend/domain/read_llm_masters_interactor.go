package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/domain/dto"
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
	"github.com/unyooon/prompt-sharing/infrastructure/repository"
)

// ReadLlmMasterInteractorはLlmMasterエンティティの読み取りを処理します。
type ReadLlmMastersInteractor struct {
	LlmMasterRepository repository.LlmMasterRepositoryInterface
}

// NewReadLlmMasterInteractorはReadLlmMasterInteractorの新しいインスタンスを作成します。
func NewReadLlmMastersInteractor(r repository.LlmMasterRepositoryInterface) *ReadLlmMastersInteractor {
	return &ReadLlmMastersInteractor{
		LlmMasterRepository: r,
	}
}

// ExecuteはLlmMasterエンティティを読み取り、レスポンスとして返します。
func (i *ReadLlmMastersInteractor) Execute(c *gin.Context) ([]dto.ReadLlmMasterResponse, *exception.CustomException) {
	llmMasters, llmMastersErr := i.LlmMasterRepository.ReadLlmsMaster()
	if llmMastersErr != nil {
		return []dto.ReadLlmMasterResponse{}, llmMastersErr
	}

	return ConvertLlmMastersToResponse(llmMasters), nil
}

// ConvertLlmMastersToResponseはLlmMasterエンティティのリストをレスポンス形式に変換します。
func ConvertLlmMastersToResponse(llmMasters []entity.LlmMaster) []dto.ReadLlmMasterResponse {
	var res []dto.ReadLlmMasterResponse
	for _, llm := range llmMasters {
		res = append(res, dto.ReadLlmMasterResponse{
			LlmId:       llm.ID,
			DisplayName: llm.ModelDisplayName,
		})
	}
	return res
}
