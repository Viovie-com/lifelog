package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Viovie-com/lifelog/internal"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var dsn string

func init() {
	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=10s&parseTime=true",
		internal.Config.DB.Account,
		internal.Config.DB.Password,
		internal.Config.DB.Host,
		internal.Config.DB.Port,
		internal.Config.DB.Database,
		internal.Config.DB.Charset)
}

func Instance() (db *gorm.DB) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("Db connection failed:" + err.Error())
	}
	return
}
