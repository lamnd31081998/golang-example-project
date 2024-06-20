package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var MasterDB *gorm.DB

func ConnectMasterDB() {
	var err error

	MasterDB, err = gorm.Open(postgres.Open(os.Getenv("MASTER_DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatalln("ConnectMasterDB Err === ", err)
	}
}
