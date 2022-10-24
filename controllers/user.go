package controllers

import (
	"fmt"
	"ilmudata/shoppingchart/database"
	"ilmudata/shoppingchart/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	//"strconv"
)
type SignIn struct {
	Email    string `form:"email" json:"email"  validate:"required"`
	Password string `form:"password"  json:"password"  validate:"required"`
}


type UserController struct {
	//declare variables
	Db    *gorm.DB
	store *session.Store
}



func InitUserController(s *session.Store) *UserController {
	db := database.DatabaseInit()
	// gorm
	db.AutoMigrate(&models.User{})

	return &UserController{Db: db, store: s}
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	return c.Render("registrasi", fiber.Map{
		"Title": "Registrasi",
	})

}

func (controller *UserController) NewRegister(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var registrasi models.User

	if err := c.BodyParser(&registrasi); err != nil {
		return c.Redirect("/login")
	}
	// save registrasi
	err := models.Registrasi(controller.Db, &registrasi)
	if err != nil {
		return c.Redirect("/login")
	}
	// if succeed
	return c.Redirect("/login")
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

func (controller *UserController) LoginUser(c *fiber.Ctx) error {

	sess, errs := controller.store.Get(c)
	if errs != nil {
		panic(errs)
	}
	var myform SignIn
	//var user models.User

	var user models.User

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}
	err := models.GetUserByEmail(controller.Db, &user, myform.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given email is not found",
		}) // http 500 internal server error
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	sHash := string(bytes)
	compareUserPassword := bcrypt.CompareHashAndPassword([]byte(sHash), []byte(user.Password))
	fmt.Println("Password Sama ?", compareUserPassword == nil)
	
	// fmt.Println("Hash Password:", user.Password)
	// fmt.Println("Hash Password2:", sHash)
	//password := compareUserPassword == nil

	// sess,err := controller.store.Get(c)
	// if err!=nil {
	// 	panic(err)
	// }

	if user.Email == myform.Email {
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password)); err != nil {

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "Password is empty",
			})

		}
		sess.Set("email", user.Email)
		sess.Set("id_user", user.IdUser)
		sess.Save()

		

		return c.Redirect("/products")

	}

	// passwordCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password))
	// if passwordCompare == nil{
	// 	sess.Set("username", user.Name)
	// 	sess.Set("userId", user.IdUser)
	// 	sess.Save()

	// 	return c.Redirect("/products")

	// }

	return c.Redirect("/login")
}
