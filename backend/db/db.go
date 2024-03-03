package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Initialize() *gorm.DB {

	if DB != nil {
		return DB
	}
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "cumulus.", // schema name
			SingularTable: false,
		}})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		panic("DB uninialized!")
	}
	return DB
}
