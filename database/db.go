package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//make connection to mysql server
func Connect() *gorm.DB {
	mysqlDSN := "root:@tcp(127.0.0.1:3306)/quizy?charset=utf8mb4&parseTime=True&loc=Local"
	db, e := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if e != nil {
		panic(e.Error())
	}

	return db
}
