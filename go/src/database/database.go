package database

import (
	"api/src/models"
)

var dbInstance = []models.People{
	{

		FirstName: "Nas",
		LastName:  "Issa",
		YearBorn:  "1997",
		Age:       24,
	},
	{

		FirstName: "John",
		LastName:  "Doe",
		YearBorn:  "1983",
		Age:       29,
	},
}

func GetData() []models.People {
	return dbInstance
}
