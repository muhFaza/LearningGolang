package models

import "time"

type Product struct {
	ID          	int64     `json:"id"`
	Name        	string    `json:"name"`
	Category    	string    `json:"category"`
	Stock       	int64    `json:"stock"`
	Price       	int64    `json:"price"`
	WarehouseID 	string    `json:"WarehouseId"`
	CreatedAt   	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}
