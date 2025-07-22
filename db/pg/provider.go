package pg

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Provide *gorm.DB for wire injection
func ProvideGormDB(db *DB) *gorm.DB {
	return db.DB
}

var ProviderSet = wire.NewSet(
	NewDB,
	ProvideGormDB, // Add this to provide *gorm.DB
)
