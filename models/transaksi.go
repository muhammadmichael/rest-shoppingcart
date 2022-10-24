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

func CreateTransaksi(db *gorm.DB, newTransaksi *Transaksi, userId uint, products []*Product) (err error) {
	newTransaksi.UserID = userId
	newTransaksi.Products = products
	err = db.Create(newTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertProductToTransaksi(db *gorm.DB, insertedTransaksi *Cart, product *Product) (err error) {
	insertedTransaksi.Products = append(insertedTransaksi.Products, product)
	err = db.Save(insertedTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadAllProductsInTransaksi(db *gorm.DB, transaksi *Transaksi, id int) (err error) {
	err = db.Where("id=?", id).Preload("Products").Find(transaksi).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadTransaksiById(db *gorm.DB, transaksis *[]Transaksi, id int) (err error) {
	err = db.Where("user_id=?", id).Find(transaksis).Error
	if err != nil {
		return err
	}
	return nil
}
