package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/unyooon/prompt-sharing/domain/setting"
	"github.com/unyooon/prompt-sharing/entity"
)

type Db struct {
	Connection *gorm.DB
}

func NewDb(setting setting.Setting) *Db {
	db, err := gorm.Open("mssql", setting.Dsn)
	if err != nil {
		panic(err.Error())
	}

	d := &Db{
		Connection: db,
	}
	d.autoMigration()

	return d
}

func (d *Db) Connect() *gorm.DB {
	return d.Connection

}

func (d *Db) autoMigration() {
	d.Connection.AutoMigrate(&entity.LlmMaster{})
	d.Connection.AutoMigrate(&entity.ParameterMaster{})
	d.Connection.AutoMigrate(&entity.UserAttributeMaster{})
	d.Connection.AutoMigrate(&entity.Liked{})
	d.Connection.AutoMigrate(&entity.User{})
	d.Connection.AutoMigrate(&entity.Parameter{})
	d.Connection.AutoMigrate(&entity.Prompt{})
	d.Connection.AutoMigrate(&entity.PromptUseHistory{})
}
