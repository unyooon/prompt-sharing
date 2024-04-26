package entity

import "github.com/jinzhu/gorm"

type LlmMaster struct {
	gorm.Model
	ModelName        string `column:"modelName"`
	ModelDisplayName string `column:"modelDisplayName"`
}
