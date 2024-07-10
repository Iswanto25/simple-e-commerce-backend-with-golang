// models/ordersModels.go

package models

import "time"

type Orders struct {
	Id          int   	`json:"idOrders" gorm:"primaryKey"`
	ProductID   string    `json:"idProducts"`
	CustomerID 	string    `json:"idCustomers"`
	Status	  	string    `json:"status" gorm:"default:pending"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}