package entity

import "time"

type User struct {
	ID                    string `gorm:"primary_key" column:"id"`
	MailAddress           string `gorm:"type:varchar(100);unique_index" column:"mailaddress"`
	DisplayName           string `gorm:"type:varchar(100)" column:"displayName"`
	UserAttributeMasterId uint   `column:"userAttributeId"`
	UserAttribute         UserAttributeMaster
	IsDeleted             bool      `column:isDeleted`
	CreatedAt             time.Time `column:createdAt`
	UpdateAt              time.Time `column:updateAt`
}
