// models/productsModels.go

package models

import "time"

type Products struct {
	Id        int       `json:"idproducts" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
	Stock     int       `json:"stock" binding:"required"`
	Status    string    `json:"status" gorm:"default:pending"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}