package db

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func InitDb() {
	var err error

	dsn := url.URL{
		User:     url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASSWWORD")),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		Path:     os.Getenv("DB_NAME"),
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&Skill{})

	if err != nil {
		panic(err)
	}

	DB = db
}

// type DbConfig struct {
// 	host     string
// 	port     int
// 	user     string
// 	dbname   string
// 	password string
// }

// var config = DbConfig{"localhost", 5432, "veritem", "test", "code"}

// func getDataBaseUrl() string {
// 	return fmt.Sprintf(
// 		"host=%s port=%d user=%s dbname=%s password=%s", config.host, config.port, config.user, config.dbname, config.password)
// }

// func GetDatabase() (*gorm.DB, error) {
// 	db, error := gorm.Open("postgres", getDataBaseUrl())
// 	return db, error
// }

// func RunMigrations(db *gorm.DB) {
//   if !db.HasTable(&model.)
// }
