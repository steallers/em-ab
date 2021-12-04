package model

import "gorm.io/gorm"

func Migration(db *gorm.DB) error {
	db.AutoMigrate()
	return nil
}
