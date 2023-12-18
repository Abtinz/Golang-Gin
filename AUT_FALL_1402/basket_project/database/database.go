package database

import (
	//postgress and gorm are needed for database
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//log will import for saving server logs in database section
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//dsn is set just same as gorm.io documentation, no need to ssl mode ...
	dsn := "host=localhost user=postgres password=abtin dbname=postgres port=5432 sslmode=disable"

	opened_database, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal("Error(database):", error)
	}

	DB = opened_database

	return DB
}

func CloseDB() {

	if DB != nil {

		sqlDB, error := DB.DB()

		if error != nil {
			log.Fatal("Error(database):", error)
		}

		sqlDB.Close()
	} else {
		log.Fatal("Error(database): no database for closing XD")

	}
}
