package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(models ...interface{}) *gorm.DB {
	dbUrl := os.Getenv("DB_URL")

	fmt.Println("dbUrl:", dbUrl)
	// fmt.Println("DB_URL=%q\n", dbUrl)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// TODO: pass the models
	if errMigrate := db.AutoMigrate(models...); errMigrate != nil {
		log.Fatal(errMigrate)
	}
	return db
}
