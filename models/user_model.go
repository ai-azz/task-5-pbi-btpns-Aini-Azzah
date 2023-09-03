package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
    Username  string         `json:"username" gorm:"not null"`
    Email     string         `json:"email" gorm:"unique;not null"`
    Password  string         `json:"-" gorm:"not null"`
    Photos    []Photo        `json:"photos" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}