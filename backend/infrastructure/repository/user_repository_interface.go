package repository

import (
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
)

type UserRepositoryInterface interface {
	Create(user entity.User) (entity.User, *exception.CustomException)
	ReadUser(query interface{}, args ...interface{}) (entity.User, *exception.CustomException)
}
