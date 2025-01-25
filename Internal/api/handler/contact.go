package handler

import (
	"fmt"
	"simple-api/Internal/models"
	"simple-api/Internal/repository/contactInfo"
	"simple-api/Internal/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetContactsList(c *fiber.Ctx) error {

	if len(contactInfo.ContactList) == 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "No contacts available",
		})
	}
	return c.JSON(contactInfo.ContactList)
}

func GetContactByName(c *fiber.Ctx) error {
	name := c.Params("name")
	name = strings.ReplaceAll(name, "%20", " ")

	_, founded := contactInfo.ContactListMap[name]
	if !founded {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Contact not found",
		})
	}

	return c.JSON(contactInfo.ContactListMap[name])
}

func AddNewContact(c *fiber.Ctx) error {

	newContact := models.Contact{}
	c.BodyParser(&newContact)

	newContactID, err := service.AddNewContact(newContact)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(map[string]string{
		"message": fmt.Sprintf("contact successfully created, ID: %d", newContactID),
	})

}

func DeleteContactByName(c *fiber.Ctx) error {
	name := c.Params("name")
	name = strings.ReplaceAll(name, "%20", " ")

	/*err:=service.DeleteContactByName(name)
	if err!=nil{
		return c.Status(fiber.statusNotFound).JSON(fiber.Map{
			"error":err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message":fmt.Sprintf("Contact %s successfully deleted",name),
	})*/
	_, founded := contactInfo.ContactListMap[name]
	if !founded {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Contact not found",
		})
	}
	delete(contactInfo.ContactListMap, name)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Contact %s successfully deleted", name),
	})
}
func UpdateContactByName(c *fiber.Ctx) error {
	name := c.Params("name")
	name = strings.ReplaceAll(name, "%20", "")
	_, found := contactInfo.ContactListMap[name]
	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Contact not found",
		})
	}
	var updates models.Contact
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	updatedContact, err := service.UpdateContact(name, updates)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Contact %s successfully updated", updatedContact.Name),
	})
}
