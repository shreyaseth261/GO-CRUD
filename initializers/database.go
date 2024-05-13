package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectToDb() {
	var err error
	//postgres://avnadmin:AVNS_SXLgo_hkok2GWW7HXnE@pg-67079ff-go-crud.l.aivencloud.com:16553/defaultdb?sslmode=require
	dsn := os.Getenv("DB_URL");
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to db", err)
	}
}