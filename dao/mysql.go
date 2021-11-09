package dao
/*
	dao层
*/

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//全局变量

var (
	DB *gorm.DB
)

func InitMySQL()(err error){
	dsn := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}

	return DB.DB().Ping()
}


func Close()  {
	DB.Close()
}

