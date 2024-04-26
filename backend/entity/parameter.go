package entity

import "github.com/jinzhu/gorm"

type Parameter struct {
	gorm.Model
	ParameterMasterId uint `column:"parameterMasterId"`
	ParameterMaster   ParameterMaster
	Value             int `column:"value"`
}
