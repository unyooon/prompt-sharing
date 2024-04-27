package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/unyooon/prompt-sharing/core/constants"
	"github.com/unyooon/prompt-sharing/domain/exception"
	"github.com/unyooon/prompt-sharing/entity"
	"github.com/unyooon/prompt-sharing/infrastructure/db"
)

type ParameterMasterRepository struct {
	db *gorm.DB
}

func NewParameterMasterRepository(db db.DbInterface) ParameterMasterRepository {
	d := db.Connect()
	return ParameterMasterRepository{
		db: d,
	}
}

// 全パラメータマスタの取得
func (repo *ParameterMasterRepository) ReadParameterMasters(query interface{}, args ...interface{}) ([]entity.ParameterMaster, *exception.CustomException) {
	var parameterMasters []entity.ParameterMaster
	if err := repo.db.Where(query, args).Find(&parameterMasters).Error; err != nil {
		return parameterMasters, &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}
	return parameterMasters, nil
}
