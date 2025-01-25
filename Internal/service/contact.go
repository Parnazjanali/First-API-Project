package service

import (
	"fmt"
	"simple-api/Internal/models"
	"simple-api/Internal/repository/contactInfo"
)

func AddNewContact(contact models.Contact) (int, error) {
	newContactID := GenerateRandomID()

	contact.ID = newContactID

	contactInfo.ContactList = append(contactInfo.ContactList, &contact)

	return newContactID, nil

}
func UpdateContact(name string, updates models.Contact) (*models.Contact, error) {
	contact, found := contactInfo.ContactListMap[name]
	if !found {
		return nil, fmt.Errorf("contact with name %s not found", name)
	}

	if updates.Name != "" {
		contact.Name = updates.Name
	}
	if updates.Phone != "" {
		contact.Phone = updates.Phone
	}
	if updates.EmailAddress != "" {
		contact.EmailAddress = updates.EmailAddress
	}

	contactInfo.ContactListMap[contact.Name] = contact

	return &contact, nil
}
