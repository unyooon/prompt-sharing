package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/core/helper"
	"github.com/unyooon/prompt-sharing/domain/dto"
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
	"github.com/unyooon/prompt-sharing/infrastructure/repository"
)

type ReadPromptsInteractor struct {
	PromptRepository repository.PromptRepositoryInterface
}

func NewReadPromptsInteractor(r repository.PromptRepositoryInterface) *ReadPromptsInteractor {
	return &ReadPromptsInteractor{
		PromptRepository: r,
	}
}

func (i *ReadPromptsInteractor) Execute(c *gin.Context) (dto.ReadPromptsResponse, *exception.CustomException) {
	var req = &dto.ReadPromptsRequest{}
	req, qErr := helper.PaginationHelper(c, req, 50)
	if qErr != nil {
		return dto.ReadPromptsResponse{}, qErr
	}

	// クエリ文字列と引数を作成
	var args []interface{}
	var query string
	if req.LlmId != nil {
		query += "llmId = ?"
		args = append(args, req.LlmId)
	}
	if req.IsPublic {
		if query != "" {
			query += " AND "
		}
		query += "isPublic = ?"
		args = append(args, req.IsPublic)
	}

	prompts, promptsErr := i.PromptRepository.ReadPrompts(req.Size, req.Page*req.Size, query, args...)
	if promptsErr != nil {
		return dto.ReadPromptsResponse{}, promptsErr
	}

	data := ConvertPromptsToResponse(prompts)

	return dto.ReadPromptsResponse{
		Page: dto.BasePageResponse{
			Number:        req.Page,
			Size:          req.Size,
			TotalElements: len(prompts),
			TotalPages:    len(prompts)/req.Size + 1,
		},
		Data: data,
	}, nil
}

func ConvertPromptsToResponse(prompts []entity.Prompt) []dto.ReadPromptResponse {
	data := []dto.ReadPromptResponse{}
	parameterResponse := []dto.ParameterResponse{}
	for _, p := range prompts {
		for _, pm := range p.Parameters {
			parameterResponse = append(parameterResponse, dto.ParameterResponse{
				Name:  pm.ParameterMaster.ParameterName,
				Value: pm.Value,
			})
		}
		data = append(data, dto.ReadPromptResponse{
			ID:          p.ID,
			LlmName:     p.Llm.ModelDisplayName,
			Parameters:  parameterResponse,
			Text:        p.Text,
			Description: p.Description,
			IsPublic:    p.IsPublic,
			CreatedUser: dto.CreatedUserResponse{
				DisplayName: p.CreatedBy,
			},
			CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return data
}
