package main

import (
	"ilmudata/shoppingchart/controllers"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	engine := html.New("./views", ".html")
	store := session.New()

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/public", "./public")

	CheckLogin := func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		val := sess.Get("email")
		if val != nil {
			return c.Next()
		}

		return c.Redirect("/login")
	}

	// INITIAL ROUTE

	//initial database
	// database.InitDb()
	userControllers := controllers.InitUserController(store)
	prodController := controllers.InitProductController(store)
	cartController := controllers.InitCartController()

	prod := app.Group("/products")
	prod.Get("/",CheckLogin, prodController.IndexProduct)
	//create
	prod.Get("/create", prodController.AddProduct)
	prod.Post("/create", prodController.AddPostedProduct)
	//edit
	prod.Get("/editproduct/:id", prodController.GetDataForEdit)
	prod.Post("/editproduct/:id", prodController.EditProduct)

	//detail
	prod.Get("/productdetail", prodController.GetDetailProduct)
	prod.Get("/detail/:id", prodController.GetDetailProduct2)

	//delete
	prod.Get("/deleteproduct/:id", prodController.DeleteProduct)

	prod.Get("/addtocart/productdetail/:productid", cartController.InsertToCart)


	//Cart
	prod.Get("/cart", cartController.GetCart)

	//Modul registrasi
	registrasi := app.Group("/register")
	registrasi.Get("/", userControllers.Register)
	registrasi.Post("/tambah", userControllers.NewRegister)

	//module login
	login := app.Group("/login")
	login.Get("/", userControllers.Login)
	login.Post("/sigIn", userControllers.LoginUser)

	//route.InitRoute(app)

	app.Listen(":3000")

}
