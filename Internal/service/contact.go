package service

import (
	"fmt"
	"simple-api/Internal/models"
	"simple-api/Internal/repository/sqliteDb"
)

func GetAllContacts() ([]models.Contact, error) {
	return sqliteDb.GetAllContacts()
}

func GetContactByID(id string) (*models.Contact, error) {
	return sqliteDb.GetContactByID(id)
}

func AddNewContact(name, phone, email string) (uint, error) {

	newContact := models.Contact{
		Name:         name,
		Phone:        phone,
		EmailAddress: email,
	}

	newContactID, err := sqliteDb.CreateContact(&newContact)
	if err != nil {
		return 0, fmt.Errorf("failed to create contact: %v", err)
	}

	return newContactID, nil
}

func DeleteContactByID(id string) error {
	exists, err := sqliteDb.ExistingContactByID(id)
	if err != nil {
		return fmt.Errorf("error checking existing contact: %v", err)
	}
	if !exists {
		return fmt.Errorf("contact with ID %s not found", id)
	}
	if err := sqliteDb.DeleteContactByID(id); err != nil {
		return fmt.Errorf("failed to delete contact with ID %s: %v", id, err)
	}
	return nil
}

func UpdateContact(id string, updatedContact *models.Contact) error {
	// بررسی وجود مخاطب با ID
	exists, err := sqliteDb.ExistingContactByID(id)
	if err != nil {
		return fmt.Errorf("error checking existing contact: %v", err)
	}
	if !exists {
		return fmt.Errorf("contact with ID %s not found", id)
	}

	// بروزرسانی اطلاعات مخاطب
	if err := sqliteDb.UpdateContact(id, updatedContact); err != nil {
		return fmt.Errorf("failed to update contact: %v", err)
	}

	return nil
}
