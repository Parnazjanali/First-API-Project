package apiserver

import (
	"simple-api/Internal/api/handler"
	"simple-api/Internal/repository/sqliteDb"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	sqliteDb.ConnectDb()
	app := fiber.New()

	app.Get("/contact", handler.GetContactsList)
	app.Get("/contact/:ID", handler.GetContactById)
	app.Post("/contact", handler.AddNewContactById)
	app.Delete("/contact/:ID", handler.DeleteContactById)
	app.Patch("/contact/:ID", handler.UpdateContactById)

	app.Listen(":8080")

}
