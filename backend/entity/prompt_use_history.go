package entity

import "github.com/jinzhu/gorm"

type PromptUseHistory struct {
	gorm.Model
	PromptId uint   `column:"promptId"`
	UserId   string `column:"userId"`
}
