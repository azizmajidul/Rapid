package models

import (
	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model
	Id       int `form:"id" json: "id" validate:"required"`
	UserID   uint
	Products []*Product `gorm:"many2many:transaksi_products;"`
}
