package store

import (
	"github.com/jinzhu/gorm"
	"github.com/yoneyan/res-mgmt/pkg/api/core/item"
	"github.com/yoneyan/res-mgmt/pkg/api/core/token"
	"github.com/yoneyan/res-mgmt/pkg/api/core/tool/config"
	_ "gorm.io/driver/sqlite"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", config.Conf.DB.Path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	result := db.AutoMigrate(&item.Item{}, &token.Token{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
