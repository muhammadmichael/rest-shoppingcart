package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id         int          `form:"id" json: "id" validate:"required"`
	Name       string       `form:"name" json: "name" validate:"required"`
	Image      string       `form:"image" json: "image" validate:"required"`
	Quantity   int          `form:"quantity" json: "quantity" validate:"required"`
	Price      float32      `form:"price" json: "price" validate:"required"`
	Carts      []*Cart      `gorm:"many2many:cart_products;"`
	Transaksis []*Transaksi `gorm:"many2many:transaksi_products;"`
}

// CRUD
func CreateProduct(db *gorm.DB, newProduct *Product) (err error) {
	err = db.Create(newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProducts(db *gorm.DB, products *[]Product) (err error) {
	err = db.Find(products).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProductById(db *gorm.DB, product *Product, id int) (err error) {
	err = db.Where("id=?", id).First(product).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *gorm.DB, product *Product) (err error) {
	db.Save(product)

	return nil
}

func DeleteProductById(db *gorm.DB, product *Product, id int) (err error) {
	db.Where("id=?", id).Delete(product)

	return nil
}
