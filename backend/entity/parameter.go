package entity

import "github.com/jinzhu/gorm"

type Parameter struct {
	gorm.Model
	PromptId          uint `column:"promptId"`
	ParameterMasterId uint `column:"parameterMasterId"`
	ParameterMaster   ParameterMaster
	Value             int `column:"value"`
}
