package persistence

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

type DB struct {
    DB    gorm.DB
}

func Start() (*DB, error) {
    db, err := gorm.Open("sqlite3", "/tmp/main.db")

    return &DB{DB: db}, err
}

func (db *DB) Close() {
    db.DB.DB().Close()
}