package contactInfo

import "simple-api/Internal/models"

var ContactList = []*models.Contact{
	{ID: 1, Name: "Parnaz Janali", Phone: "11111111", EmailAddress: "example@gmail.com"},
	{ID: 2, Name: "Saba Saberi", Phone: "55555", EmailAddress: "example@gmail.com"},
	{ID: 3, Name: "H Hamedi", Phone: "44444", EmailAddress: "example@gmail.com"},
}

var ContactListMap = map[string]models.Contact{
	"Parnaz Janali": models.Contact{
		Name:         "Parnaz Janali",
		Phone:        "11111111",
		EmailAddress: "m@g.com",
	},
	"Saba Saberi": models.Contact{
		Name:         "Saba Saberi",
		Phone:        "55555",
		EmailAddress: "s@g.com",
	},
	"H Hamedi": models.Contact{
		Name:         "H Hamedi",
		Phone:        "44444",
		EmailAddress: "H@g.com",
	},
}
