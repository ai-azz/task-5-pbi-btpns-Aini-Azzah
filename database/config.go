package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConfigureDatabase() (*gorm.DB, error) {
    dsn := "root:@tcp(localhost:3306)/task-5-pbi?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    DB = db

    return db, nil
}


func GetDB() *gorm.DB {
    return DB
}