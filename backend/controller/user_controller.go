package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/usecase"
)

type UserController struct {
	UserCreater usecase.UserCreater
}

func NewUserController(
	userCreater usecase.UserCreater,
) *UserController {
	return &UserController{
		UserCreater: userCreater,
	}
}

// CreateUser godoc
// @Summary ユーザー登録
// @Description ユーザーを登録します
// @Tags user
// @Accept json
// @Produce json
// @Success 200
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	err := uc.UserCreater.Execute(c)
	Handler(c, nil, err)
}
