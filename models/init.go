package models

import "Hulk/db"

func init() {
	db.Db.AutoMigrate(&InterfaceConfigModel{})
}
