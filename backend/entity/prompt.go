package entity

import "github.com/jinzhu/gorm"

type Prompt struct {
	gorm.Model
	LlmId       uint `column:"llmId"`
	Llm         LlmMaster
	Parameters  []Parameter `gorm:"many2many:prompt_parameters;"`
	Text        string      `column:"text"`
	Description string      `column:"description"`
	IsPublic    bool        `column:"isPublic"`
	CreatedBy   string      `column:"createdBy"`
}
