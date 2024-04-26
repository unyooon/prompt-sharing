package entity

import "github.com/jinzhu/gorm"

type UserAttributeMaster struct {
	gorm.Model
	AttributeName string `column:"attributeName"`
}
