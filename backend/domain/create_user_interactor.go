package domain

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/core/constants"
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
	"github.com/unyooon/prompt-sharing/infrastructure/repository"
)

type CreateUserInteractor struct {
	UserRepository repository.UserRepositoryInterface
}

func NewCreateUserInteractor(userRepository repository.UserRepositoryInterface) *CreateUserInteractor {
	return &CreateUserInteractor{
		UserRepository: userRepository,
	}
}

func (i *CreateUserInteractor) Execute(c *gin.Context) *exception.CustomException {
	uid, existsUid := c.Get("userId")
	mailAddress, existsMail := c.Get("mailAddress")
	if !existsUid || !existsMail {
		e := &exception.CustomException{
			Code:    constants.NotFoundCode,
			Message: constants.NotFoundUserId,
			Err:     errors.New("not found userId or mailAddress"),
		}
		return e
	}

	user := entity.User{
		ID:                    uid.(string),
		MailAddress:           mailAddress.(string),
		DisplayName:           "",
		UserAttributeMasterId: 3,
	}

	// ユーザ作成
	createdUser, err := i.UserRepository.Create(user)
	if err != nil {
		return err
	}

	// レスポンス
	c.JSON(http.StatusOK, createdUser)
	return nil
}
