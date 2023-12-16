package models

import "time"

type Product struct {
	ID          	int64     `json:"id"`
	Name        	string    `json:"name"`
	Category    	string    `json:"category"`
	Stock       	string    `json:"stock"`
	Price       	string    `json:"price"`
	WarehouseID 	string    `json:"WarehouseId"`
	CreatedAt   	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}
