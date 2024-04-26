package entity

import "github.com/jinzhu/gorm"

type ParameterMaster struct {
	gorm.Model
	LlmMasterId   uint `column:"llmModelMasterId"`
	LlmMaster     LlmMaster
	ParameterName string `column:"parameterName"`
}
