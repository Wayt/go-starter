package models

import (
	"time"
)

// BaseVariables visible variable
var BaseVariables = []string{"id", "created_at", "updated_at"}

// Model object interface
type Model interface {
	TableName() string
}

// Base model structure
type Base struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"Updated_at"`
}
