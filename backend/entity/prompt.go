package entity

import "github.com/jinzhu/gorm"

type Prompt struct {
	gorm.Model
	Text        string `column:"text"`
	Description string `column:"description"`
	IsPublic    bool   `column:"isPublic"`
	CreatedBy   string `column:"createdBy"`
}
