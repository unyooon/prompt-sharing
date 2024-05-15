package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/unyooon/prompt-sharing/core/constants"
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
	"github.com/unyooon/prompt-sharing/infrastructure/db"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db db.DbInterface) *UserRepository {
	d := db.Connect()
	return &UserRepository{
		db: d,
	}
}

func (repo *UserRepository) ReadUser(query interface{}, args ...interface{}) (entity.User, *exception.CustomException) {
	var users entity.User
	if err := repo.db.Where(query, args).First(&users).Error; err != nil {
		return users, &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}
	return users, nil
}

func (repo *UserRepository) Create(user entity.User) (entity.User, *exception.CustomException) {
	if err := repo.db.Create(&user).Error; err != nil {
		return user, &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}
	return user, nil
}
