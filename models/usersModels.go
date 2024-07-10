// models/usersModels.go

package models

import "time"

type Users struct {
    ID        string    `json:"id" gorm:"type:varchar(15);primaryKey"`
    Photos    string    `json:"photos"`
    Name      string    `json:"name"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
