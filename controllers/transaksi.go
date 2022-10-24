package controllers

import (
	"rapid/rest-shoppingcart/database"
	"rapid/rest-shoppingcart/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransaksiController struct {
	// Declare variables
	Db *gorm.DB
}

func InitTransaksiController() *TransaksiController {
	db := database.InitDb()
	// gorm sync
	db.AutoMigrate(&models.Transaksi{})

	return &TransaksiController{Db: db}
}

// GET /checkout/:userid
func (controller *TransaksiController) InsertToTransaksi(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intUserId, _ := strconv.Atoi(params["userid"])

	var transaksi models.Transaksi
	var cart models.Cart

	// Find the cart
	errNoCart := models.ReadCartById(controller.Db, &cart, intUserId)
	if errNoCart != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Tidak dapat menemukan Cart dengan Id " + params["userid"] + ", Gagal Melakukan Checkout",
		})
	}

	// Find the product first,
	err := models.ReadAllProductsInCart(controller.Db, &cart, intUserId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// Jika Cart kosong
	if len(cart.Products) == 0 {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Cart kosong, silahkan isi Product ke Cart terlebih dahulu",
		})
	}

	errs := models.CreateTransaksi(controller.Db, &transaksi, uint(intUserId), cart.Products)
	if errs != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// Delete products in cart
	errss := models.UpdateCart(controller.Db, cart.Products, &cart, uint(intUserId))
	if errss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// if succeed
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil Melakukan Checkout",
	})
}

// GET /history/:userid
func (controller *TransaksiController) GetTransaksi(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intUserId, _ := strconv.Atoi(params["userid"])

	var transaksis []models.Transaksi
	err := models.ReadTransaksiById(controller.Db, &transaksis, intUserId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Message":    "History Transaksi",
		"Transaksis": transaksis,
	})

}

// GET /history/detail/:transaksiid
func (controller *TransaksiController) DetailTransaksi(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intTransaksiId, _ := strconv.Atoi(params["transaksiid"])

	var transaksi models.Transaksi
	err := models.ReadAllProductsInTransaksi(controller.Db, &transaksi, intTransaksiId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Message":  "Detail Product pada Transaksi",
		"Products": transaksi.Products,
	})
}
