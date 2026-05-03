package models

import "time"

type User struct {
	ID         string     `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	Email      string     `json:"email" db:"email"`
	Address    string     `json:"address" db:"address"`
	Role       string     `json:"role" db:"role" `
	CreatedAt  time.Time  `json:"createdAt" db:"Created_at"`
	ArchivedAt *time.Time `json:"archivedAt" db:"archived_at"`
}

type Product struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Price     float64   `json:"price" db:"price"`
	Image     string    `json:"image" db:"image"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Order struct {
	ID              string    `json:"id" db:"id"`
	UID             string    `json:"uid" db:"user_id"`
	TotalAmount     float64   `json:"totalAmount" db:"total_amount"`
	Status          string    `json:"status" db:"status"`
	ShippingAddress string    `json:"shippingAddress" db:"shipping_address"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}

type OrderItems struct {
	ID         string    `json:"id" db:"id"`
	OID        string    `json:"oid" db:"order_id"`
	PID        string    `json:"productID" db:"product_id"`
	OrderPrice float64   `json:"orderPrice" db:"order_price"`
	Quantity   int       `json:"quantity" db:"quantity"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

type Session struct {
	ID        string    `json:"id" db:"id"`
	UID       string    `json:"uid" db:"user_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
