package model

import (
	"time"
)

type Customer struct {
	CustomerID  int       `gorm:"column:customer_id;AUTO_INCREMENT;primary_key"`
	FirstName   string    `gorm:"column:first_name;NOT NULL"`
	LastName    string    `gorm:"column:last_name;NOT NULL"`
	Email       string    `gorm:"column:email;unique;NOT NULL"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Address     string    `gorm:"column:address"`
	Password    string    `gorm:"column:password;NOT NULL"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
}

func (m *Customer) TableName() string {
	return "customer"
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateAddressRequest struct {
	Address string `json:"address" binding:"required"`
}

type LoginResponse struct {
	Customer Customer `json:"customer"`
	Token    string   `json:"token"`
}
