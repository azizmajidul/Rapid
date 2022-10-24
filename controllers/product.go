package controllers

import (
	"fmt"
	"ilmudata/shoppingchart/database"
	"ilmudata/shoppingchart/models"
	"strconv"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductController struct {
	// declare variables
	Db *gorm.DB
	store *session.Store
}

func InitProductController(s *session.Store) *ProductController {
	db := database.DatabaseInit()
	// gorm
	db.AutoMigrate(&models.Product{})

	return &ProductController{Db: db, store: s}
}

// routing
// GET /products
func (controller *ProductController) IndexProduct(c *fiber.Ctx) error {
	// load all products
	var products []models.Product
	

	err := models.ReadProduct(controller.Db, &products)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	val := sess.Get("id_user")
	return c.Render("products", fiber.Map{
		"Title":    "Daftar Produk",
		"Products": products,
		"UserId":   val,
	})
}

// GET /products/create
func (controller *ProductController) AddProduct(c *fiber.Ctx) error {
	return c.Render("addProduct", fiber.Map{
		"Title": "Tambah Produk",
	})
}

// POST /products/create
func (controller *ProductController) AddPostedProduct(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var myform models.Product

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}

	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "documents" key:
		files := form.File["image"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			if err := c.SaveFile(file, fmt.Sprintf("./upload/%s", file.Filename)); err != nil {
				return err
			}

		}

	}
	// save product
	err := models.CreateProduct(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/products")
	}
	// if succeed
	return c.Redirect("/products")
}

func (controller *ProductController) GetDetailProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}

// GET /products/detail/xxx
func (controller *ProductController) GetDetailProduct2(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}

//Get Product to Edit

func (controller *ProductController) GetDataForEdit(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("editproduct", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}

func (controller *ProductController) EditProduct(c *fiber.Ctx) error {
	//myform := new(models.Product)
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var myform models.Product

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}
	product.Name = myform.Name
	product.Quantity = myform.Quantity
	product.Price = myform.Price
	// save product
	models.UpdateProduct(controller.Db, &product)

	return c.Redirect("/products")
}

func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	models.DeleteProductById(controller.Db, &product, idn)
	return c.Redirect("/products")
}
