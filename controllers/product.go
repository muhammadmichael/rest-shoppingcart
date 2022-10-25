package controllers

import (
	"fmt"
	"rapid/rest-shoppingcart/database"
	"rapid/rest-shoppingcart/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductController struct {
	// Declare variables
	Db *gorm.DB
}

func InitProductController() *ProductController {
	db := database.InitDb()
	// gorm sync
	db.AutoMigrate(&models.Product{})

	return &ProductController{Db: db}
}

// @BasePath /

// GetAllProduct godoc
// @Summary GetAllProduct example
// @Schemes
// @Description GetAllProduct
// @Tags rest-shoppingcart
// @Accept json
// @Produce json
// @Success 200 {json} GetAllProduct
// @Router /products [get]
func (controller *ProductController) GetAllProduct(c *fiber.Ctx) error {
	// Load all Products
	var products []models.Product
	err := models.ReadProducts(controller.Db, &products)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.JSON(fiber.Map{
		"Message":  "Berhasil mendapatkan seluruh list products",
		"Products": products,
	})
}

// @BasePath /

// AddPostedProduct godoc
// @Summary AddPostedProduct example
// @Schemes
// @Description AddPostedProduct
// @Tags rest-shoppingcart
// @Param name formData string true "Product Name"
// @Param quantity formData int true "Quantity"
// @Param price formData number true "Price"
// @Param image formData file true "Image"
// @Accept mpfd
// @Produce json
// @Success 200 {json} AddPostedProduct
// @Security ApiKeyAuth
// @Router /products/create [post]
func (controller *ProductController) AddPostedProduct(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request, Product Form is not complete",
		})
	}

	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "image" key:
		files := form.File["image"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			product.Image = fmt.Sprintf("public/upload/%s", file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("public/upload/%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	// save product
	err := models.CreateProduct(controller.Db, &product)
	if err != nil {
		return c.SendStatus(500)
	}
	// if succeed
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil Menambahkan Product",
	})
}

// @BasePath /

// DetailProduct godoc
// @Summary DetailProduct example
// @Schemes
// @Description DetailProduct
// @Param        id         path      int     true  "Product Id"       minimum(1)
// @Tags rest-shoppingcart
// @Accept json
// @Produce json
// @Success 200 {json} DetailProduct
// @Router /products/detail/{id} [get]
func (controller *ProductController) DetailProduct(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intId, errs := strconv.Atoi(params["id"])

	if errs != nil {
		fmt.Println(errs)
	}

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, intId)
	if err != nil {
		return c.JSON(fiber.Map{
			"Status":  500,
			"message": "Tidak dapat mencari product dengan Id " + params["id"],
		})
	}

	return c.JSON(fiber.Map{
		"message": "Detail Product Dengan Id " + params["id"],
		"product": product,
	})
}

// @BasePath /

// AddUpdatedProduct godoc
// @Summary AddUpdatedProduct example
// @Schemes
// @Description AddUpdatedProduct
// @Tags rest-shoppingcart
// @Param        id         path      int     true  "Product Id"       minimum(1)
// @Param name formData string true "Product Name"
// @Param quantity formData int true "Quantity"
// @Param price formData number true "Price"
// @Param image formData file true "Image"
// @Accept mpfd
// @Produce json
// @Success 200 {json} AddUpdatedProduct
// @Security ApiKeyAuth
// @Router /products/ubah/{id} [put]
func (controller *ProductController) AddUpdatedProduct(c *fiber.Ctx) error {
	var product models.Product
	var checkProduct models.Product

	params := c.AllParams() // "{"id": "1"}"
	intId, _ := strconv.Atoi(params["id"])
	product.Id = intId

	if err := c.BodyParser(&product); err != nil {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request, Product Form is not complete",
		})
	}

	errFind := models.ReadProductById(controller.Db, &checkProduct, intId)
	if errFind != nil {
		return c.JSON(fiber.Map{
			"Status":  500,
			"message": "Tidak dapat mencari product dengan Id " + params["id"] + ", Gagal Mengupdate Product",
		})
	}

	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "image" key:
		files := form.File["image"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			product.Image = fmt.Sprintf("public/upload/%s", file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("public/upload/%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	// save product
	err := models.UpdateProduct(controller.Db, &product)
	if err != nil {
		return c.SendStatus(500)
	}
	// if succeed
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil Mengubah Product dengan Id " + params["id"],
	})
}

// @BasePath /api

// DeleteProduct godoc
// @Summary DeleteProduct example
// @Schemes
// @Description DeleteProduct
// @Param        id         path      int     true  "Product Id"       minimum(1)
// @Tags rest-shoppingcart
// @Accept json
// @Produce json
// @Success 200 {json} DeleteProduct
// @Security ApiKeyAuth
// @Router /products/hapus/{id} [delete]
func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intId, errs := strconv.Atoi(params["id"])

	if errs != nil {
		fmt.Println(errs)
	}

	var product models.Product
	errFind := models.ReadProductById(controller.Db, &product, intId)
	if errFind != nil {
		return c.JSON(fiber.Map{
			"Status":  500,
			"message": "Tidak dapat mencari product dengan Id " + params["id"] + ", Gagal Menghapus Product",
		})
	}

	err := models.DeleteProductById(controller.Db, &product, intId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// if succeed
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil Menghapus Product dengan Id " + params["id"],
	})
}
