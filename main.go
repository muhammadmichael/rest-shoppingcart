package main

import (
	"rapid/rest-shoppingcart/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// static
	app.Static("/", "./public", fiber.Static{
		Index: "",
	})

	// controllers
	prodController := controllers.InitProductController()
	authController := controllers.InitAuthController()
	cartController := controllers.InitCartController()
	transaksiController := controllers.InitTransaksiController()

	prod := app.Group("/products")
	prod.Get("/", prodController.GetAllProduct)
	prod.Get("/create", prodController.AddProduct)
	prod.Post("/create", prodController.AddPostedProduct)
	prod.Get("/detail/:id", prodController.DetailProduct)
	prod.Get("/ubah/:id", prodController.UpdateProduct)
	prod.Post("/ubah/:id", prodController.AddUpdatedProduct)
	prod.Get("/hapus/:id", prodController.DeleteProduct)
	prod.Get("/addtocart/:cartid/product/:productid", cartController.InsertToCart)

	cart := app.Group("/shoppingcart")
	cart.Get("/:cartid", cartController.GetShoppingCart)

	transaksi := app.Group("/checkout")
	transaksi.Get("/:userid", transaksiController.InsertToTransaksi)

	history := app.Group("/history")
	history.Get("/:userid", transaksiController.GetTransaksi)
	history.Get("/detail/:transaksiid", transaksiController.DetailTransaksi)

	app.Get("/login", authController.Login)
	app.Post("/login", authController.LoginPosted)
	app.Get("/register", authController.Register)
	app.Post("/register", authController.AddRegisteredUser)

	app.Listen(":3000")
}
