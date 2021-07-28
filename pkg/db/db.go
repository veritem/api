package db

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Model struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (model *Model) BeforeCreate(tx *gorm.DB) (err error) {
	model.ID = uuid.NewString()
	return
}

func Connect() error {
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
		return err
	}

	err = db.AutoMigrate(&SkillsCategory{}, &Skill{}, &Secret{}, &ProjectEcosystem{}, &Project{}, &Experience{})

	if err != nil {
		return err
	}

	DB = db

	return nil
}
