package entity

import "github.com/jinzhu/gorm"

type Liked struct {
	gorm.Model
	PromptId uint   `column:"promptId"`
	UserId   string `column:"userId"`
}
