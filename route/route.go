package route

import (
	"ilmudata/shoppingchart/controllers"

	"github.com/gofiber/fiber/v2"
)
func InitRoute(r *fiber.App) {
	r.Get("/user", controllers.UserRead )
 
}