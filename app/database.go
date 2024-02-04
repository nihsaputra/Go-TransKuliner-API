package app

import (
	"TransKuliner/halper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnectionMysql() *gorm.DB {
	host := "localhost"
	port := "3306"
	username := "root"
	password := "@Saputra03"
	dbName := "TransKuliner_DB"
	dsn := username + ":" + password + "@tcp" + "(" + host + ":" + port + ")" + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	halper.PanicIfError(err)
	return db
}
