package models

import ()

// Pony DB object
type Pony struct {
	Base
	Name  string `gorm:"type:varchar(255);not null;unique"`
	Title string `gorm:"type:varchar(2048)"`
}

// TableName for gorm
func (p Pony) TableName() string {
	return "ponies"
}

// Ponies Pony array
type Ponies []*Pony
