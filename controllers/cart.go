package controllers

import (
	"rapid/rest-shoppingcart/database"
	"rapid/rest-shoppingcart/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CartController struct {
	// Declare variables
	Db *gorm.DB
}

func InitCartController() *CartController {
	db := database.InitDb()
	// gorm sync
	db.AutoMigrate(&models.Cart{})

	return &CartController{Db: db}
}

// @BasePath /

// InsertToCart godoc
// @Summary InsertToCart example
// @Schemes
// @Description InsertToCart
// @Tags rest-shoppingcart
// @Param        cartid         path      int     true  "Cart Id"       minimum(1)
// @Param        productid         path      int     true  "Product Id"       minimum(1)
// @Accept json
// @Produce json
// @Success 200 {json} InsertToCart
// @Security ApiKeyAuth
// @Router /products/addtocart/{cartid}/product/{productid} [get]
func (controller *CartController) InsertToCart(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intCartId, _ := strconv.Atoi(params["cartid"])
	intProductId, _ := strconv.Atoi(params["productid"])

	var cart models.Cart
	var product models.Product

	// Find the product first,
	err := models.ReadProductById(controller.Db, &product, intProductId)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Tidak dapat menemukan Product dengan Id " + params["productid"] + ", Gagal menambahkan ke Shopping Cart ",
		})
	}

	// Then find the cart
	errs := models.ReadCartById(controller.Db, &cart, intCartId)
	if errs != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Tidak dapat menemukan Cart dengan Id " + params["cartid"] + ", Gagal menambahkan ke Shopping Cart ",
		})
	}

	// Finally, insert the product to cart
	errss := models.InsertProductToCart(controller.Db, &cart, &product)
	if errss != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error, Gagal menambahkan ke Shopping Cart ",
		})
	}

	// if succeed
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil Menambahkan Product dengan Id " + params["productid"] + " ke Shopping Cart " + params["cartid"],
	})
}

// @BasePath /

// GetShoppingCart godoc
// @Summary GetShoppingCart example
// @Schemes
// @Description GetShoppingCart
// @Tags rest-shoppingcart
// @Param        cartid         path      int     true  "Cart Id"       minimum(1)
// @Accept json
// @Produce json
// @Success 200 {json} GetShoppingCart
// @Security ApiKeyAuth
// @Router /shoppingcart/{cartid} [get]
func (controller *CartController) GetShoppingCart(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intCartId, _ := strconv.Atoi(params["cartid"])

	var cart models.Cart

	// Then find the cart
	errs := models.ReadCartById(controller.Db, &cart, intCartId)
	if errs != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Tidak dapat menemukan Cart dengan Id " + params["cartid"],
		})
	}

	err := models.ReadAllProductsInCart(controller.Db, &cart, intCartId)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error, Gagal mendapatkan Shopping Cart ",
		})
	}

	return c.JSON(fiber.Map{
		"Message":  "Shopping Cart dengan Id " + params["cartid"],
		"Products": cart.Products,
	})
}
