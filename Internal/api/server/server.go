package apiserver

import (
	"simple-api/Internal/api/handler"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	app.Get("/contact", handler.GetContactsList)
	app.Get("/contact/:name", handler.GetContactByName)
	app.Post("/contact", handler.AddNewContact)
	app.Delete("/contact/:name", handler.DeleteContactByName)
	app.Patch("/contact/:name", handler.UpdateContactByName)

	app.Listen(":8080")

}
