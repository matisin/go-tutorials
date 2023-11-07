package cities

import (
	"database-api/domains/geographic/entities"

	"database-api/db"
)

func CreateService(data entities.City) (entities.City, error) {

	db := db.GetDBInstance()
	result := db.Create(&data)

	if result.Error != nil {
		return entities.City{}, result.Error
	}

	return data, nil
}

func FindOneService(ID string) (entities.City, error) {
	var City entities.City

	db := db.GetDBInstance()
	result := db.First(&City, "id = ?", ID)

	if result.Error != nil {
		return entities.City{}, result.Error
	}

	return City, nil
}

func FindAllService() ([]entities.City, error) {
	var cities []entities.City
	db := db.GetDBInstance()
	result := db.Find(&cities)
	if result.Error != nil {
		return nil, result.Error
	}

	return cities, nil
}
