package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func init() {
	dsn := "host=10.2.0.101 user=zcw password=123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
