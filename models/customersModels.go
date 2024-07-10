// models/customersModels.go

package models

import "time"

type Customers struct {
	ID        string    `json:"id" gorm:"type:varchar(15);primaryKey"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}