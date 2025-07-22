package pg

import (
	"framework/domain"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewDB() (*DB, error) {
	db, err := gorm.Open(postgres.Open(viper.GetString("postgres.dsn")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// auto migrate
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}
