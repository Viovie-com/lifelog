package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Viovie-com/lifelog/internal"
)

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
		fmt.Println("Db connection failed: ", err)
	}
	return
}
