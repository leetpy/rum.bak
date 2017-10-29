package model

import "github.com/jinzhu/gorm"

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		newDb, err := newDB()
		if err != nil {
			panic(err)
		}

		db = newDb
	}
	return db
}

func newDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "rum.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
