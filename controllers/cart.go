package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ilmudata/shoppingchart/database"
	"ilmudata/shoppingchart/models"
	"strconv"



)

type CartController struct {
	// Declare variables
	Db    *gorm.DB
	//store *session.Store
}

func InitCartController() *CartController {
	db := database.DatabaseInit()
	// gorm sync
	db.AutoMigrate(&models.Cart{})

	return &CartController{Db: db, }
}

func (controller *CartController) InsertToCart(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	//intCartId, _ := strconv.Atoi(params["cartid"])
	intProductId, _ := strconv.Atoi(params["productid"])

	var cart models.Cart
	var product models.Product

	// Find the product first,
	err := models.ReadProductById(controller.Db, &product, intProductId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// // Then find the cart
	// errs := models.ReadCartById(controller.Db, &cart, intCartId)
	// if errs != nil {
	// 	return c.SendStatus(500) // http 500 internal server error
	// }

	// Finally, insert the product to cart
	errss := models.InsertProductToCart(controller.Db, &cart, &product)
	if errss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.Redirect("/products/cart")
}

func (controller *CartController) GetCart(c *fiber.Ctx) error{
	var cart []models.Cart
	//var product []models.Product

	// param := c.AllParams()
	// intCartId, _ := strconv.Atoi(param["cartid"])

	err := models.ReadCart(controller.Db, &cart)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
		
	}
	return c.Render("shoppingcart", fiber.Map{
		"Title":    "Detail Product",
		"Products": cart,
		
	})
	// return c.JSON(fiber.Map{
	// 	"Mesaage" : "Berhasil",
	// 	"Hasil" : cart, 
	// })

}





