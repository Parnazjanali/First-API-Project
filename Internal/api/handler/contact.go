package handler

import (
	"fmt"
	"simple-api/Internal/models"
	"simple-api/Internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetContactsList(c *fiber.Ctx) error {

	contacts, err := service.GetAllContacts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(contacts)
}

func GetContactById(c *fiber.Ctx) error {
	id := c.Params("id")
	contact, err := service.GetContactByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Contact not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(contact)
}

func AddNewContactById(c *fiber.Ctx) error {
	var newContact models.Contact
	if err := c.BodyParser(&newContact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Contact has already exist",
		})
	}

	newContactID, err := service.AddNewContact(newContact.Name, newContact.Phone, newContact.EmailAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Contact successfully created, ID: %d", newContactID),
	})
}

func DeleteContactById(c *fiber.Ctx) error {
	id := c.Params("id")
	err := service.DeleteContactByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]string{
			"error": "Contact not found",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(map[string]string{
		"message": "Contact successfully deleted",
	})
}

func UpdateContactById(c *fiber.Ctx) error {
	id := c.Params("id")
	updatedContact := new(models.Contact)
	if err := c.BodyParser(updatedContact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid input",
		})
	}
	err := service.UpdateContact(id, updatedContact)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]string{
			"error": "Contact not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(map[string]string{
		"message": "Contact successfully updated",
	})
}
