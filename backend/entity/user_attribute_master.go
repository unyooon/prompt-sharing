package entity

import "github.com/jinzhu/gorm"

type UserAttributeMaster struct {
	gorm.Model
	AttributeName string `gorm:"type:nvarchar(100)" column:"attributeName"`
}
