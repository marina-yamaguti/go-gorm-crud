package models

import (
	"time"
)

type Product struct {
	ID        uint       `gorm:"primaryKey" json:"id" example:"1"`
	CreatedAt time.Time  `gorm:"index" json:"-"`
	UpdatedAt time.Time  `gorm:"index" json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Name      string     `json:"name" example:"Smartphone" binding:"required"`
	Code      string     `json:"code" example:"PHN-001" binding:"required"`
	Price     float64    `json:"price" example:"699.99" binding:"required"`
}
