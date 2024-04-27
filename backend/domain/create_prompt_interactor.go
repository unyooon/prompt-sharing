package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/core/constants"
	"github.com/unyooon/prompt-sharing/core/validatation"
	"github.com/unyooon/prompt-sharing/domain/dto"
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
	"github.com/unyooon/prompt-sharing/infrastructure/repository"
)

type CreatePromptInteractor struct {
	PromptRepository          repository.PromptRepositoryInterface
	LlmMasterRepository       repository.LlmMasterRepositoryInterface
	ParameterMasterRepository repository.ParameterMasterRepositoryInterface
}

func NewCreatePromptInteractor(pr repository.PromptRepositoryInterface, lmr repository.LlmMasterRepositoryInterface, pmr repository.ParameterMasterRepositoryInterface) *CreatePromptInteractor {
	return &CreatePromptInteractor{
		PromptRepository:          pr,
		LlmMasterRepository:       lmr,
		ParameterMasterRepository: pmr,
	}
}

func (i *CreatePromptInteractor) Execute(c *gin.Context) *exception.CustomException {
	uid, exists := c.Get("userId")
	if !exists {
		e := &exception.CustomException{
			Code:    constants.NotFoundCode,
			Message: constants.NotFoundUserId,
			Err:     errors.New("not found userId"),
		}
		return e
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	req := new(dto.CreatePromptRequest)
	if ve := validatation.RequestValidate(req, c); ve != nil {
		return ve
	}

	json.Unmarshal(body, &req)

	// パラメータの定義
	var parameters []entity.Parameter
	for _, p := range req.Parameters {
		parameter := entity.Parameter{
			ParameterMasterId: p.ParameterMasterId,
			Value:             p.Value,
		}
		parameters = append(parameters, parameter)
	}

	// プロンプトの作成
	prompt := entity.Prompt{
		LlmId:       req.LlmId,
		Parameters:  parameters,
		Text:        req.Text,
		Description: req.Description,
		IsPublic:    false,
		CreatedBy:   uid.(string),
	}

	_, err := i.PromptRepository.Create(prompt)

	return err
}
