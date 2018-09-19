package db

import (
	"github.com/jinzhu/gorm"
	// driver for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User stands for users sql table
type User struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// Init inits db
func Init() (*gorm.DB, error) {
	url := "localhost"

	db, err := gorm.Open("mysql", "root:karabiner@tcp("+url+")/karabiner?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	return db, nil
}
