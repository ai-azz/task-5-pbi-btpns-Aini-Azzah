package app

import "time"

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" binding:"required"`
    Email     string    `json:"email" binding:"required,unique"`
    Password  string    `json:"password" binding:"required,min=6"`
    Photos    []Photo   `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}