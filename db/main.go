package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
}

var config = DbConfig{"localhost", 5432, "veritem", "test", "code"}

func getDataBaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s", config.host, config.port, config.user, config.dbname, config.password)
}

func GetDatabase() (*gorm.DB, error) {
	db, error := gorm.Open("postgres", getDataBaseUrl())
	return db, error
}

// func RunMigrations(db *gorm.DB) {
//   if !db.HasTable(&model.)
// }
