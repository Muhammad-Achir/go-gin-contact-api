package models

import "time"

type Contact struct {
	ID        int
	Name      string
	Phone     string
	Email     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
