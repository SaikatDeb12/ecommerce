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
	Price     int       `json:"address" db:"address"`
	Image     string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"createdAt" db:"Created_at"`
}
