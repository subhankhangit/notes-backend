package db

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kamran0812/notes-backend/conf"
	"github.com/kamran0812/notes-backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Conn *gorm.DB
)

func ConnectDB() {
	var err error
	p := conf.GetConfig("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("error: failed to get port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.GetConfig("DB_HOST"), port, conf.GetConfig("DB_USER"), conf.GetConfig("DB_PASSWORD"), conf.GetConfig("DB_NAME"))
	Conn, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("info: connected to database")
	Conn.AutoMigrate(&model.Note{})
	Conn.AutoMigrate(&model.User{})
	log.Println("info: database migrated")
}
