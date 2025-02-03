package sqliteDb

import (
	"errors"
	"fmt"
	"simple-api/Internal/models"

	"gorm.io/gorm"
)

func ExistingContactByID(id string) (bool, error) {
	var contact *models.Contact
	err := Db.Where("id = ?", id).First(&contact).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return err == nil, err
}

func GetAllContacts() ([]models.Contact, error) {
	var contact []models.Contact

	if err := Db.Find(&contact).Error; err != nil {
		return nil, err
	}

	if len(contact) == 0 {
		return nil, fmt.Errorf("no contact found")
	}

	return contact, nil
}

func GetContactByID(id string) (*models.Contact, error) {
	var contact *models.Contact
	err := Db.First(&contact, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("contact with ID %s not found", id)
		}
		return nil, err
	}
	return contact, nil
}

func DeleteContactByID(id string) error {
	result := Db.Delete(&models.Contact{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete contact with ID %s: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("contact with ID %s not found", id)
	}
	return nil
}

func CreateContact(contact *models.Contact) (uint, error) {
	if err := Db.Create(&contact).Error; err != nil {
		return 0, err
	}

	return contact.ID, nil
}

func UpdateContact(id string, updatedContact *models.Contact) error {
	result := Db.Model(&models.Contact{}).Where("id = ?", id).Updates(updatedContact)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("contact with ID %s not found", id)
	}
	return nil
}
