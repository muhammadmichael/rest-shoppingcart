package main

import (
	"rapid/rest-shoppingcart/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
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

	// Auth Routes (Register and Login)
	app.Post("/login", authController.LoginPosted)
	app.Post("/register", authController.AddRegisteredUser)

	// Unauthenticated Routes
	prod := app.Group("/products")
	prod.Get("/", prodController.GetAllProduct)
	prod.Get("/detail/:id", prodController.DetailProduct)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mysecretpassword"),
	}))

	// All the routes below need authentication
	// Product Routes (CRUD Products, Add Product to Shopping Cart)
	prod.Post("/create", prodController.AddPostedProduct)
	prod.Put("/ubah/:id", prodController.AddUpdatedProduct)
	prod.Delete("/hapus/:id", prodController.DeleteProduct)
	prod.Get("/addtocart/:cartid/product/:productid", cartController.InsertToCart)

	// Cart Routes (View Shopping Cart)
	cart := app.Group("/shoppingcart")
	cart.Get("/:cartid", cartController.GetShoppingCart)

	// Transaksi Routes (Checkout to insert Products in Shopping Cart to History Transaksi)
	transaksi := app.Group("/checkout")
	transaksi.Get("/:userid", transaksiController.InsertToTransaksi)

	// History Routes (View History Transaksi and Detail Transaksi (Detail = List of Products))
	history := app.Group("/history")
	history.Get("/:userid", transaksiController.GetTransaksi)
	history.Get("/detail/:transaksiid", transaksiController.DetailTransaksi)

	app.Listen(":3000")
}
