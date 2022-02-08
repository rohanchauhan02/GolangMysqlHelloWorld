package storage

import (
	grom "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	config "test/config"
)

var DB *grom.DB

func NewDB(params ...string) *grom.DB {
	var err error
	conString := config.GetMySQLConnectionString()
	log.Print(conString)
	DB, err = grom.Open(config.GetDBType(), conString)
	if err != nil {
		log.Panic(err)
	}
	return DB
}

func GetDbInstance() *grom.DB{
	return DB
}
