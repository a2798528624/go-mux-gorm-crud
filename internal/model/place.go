package model

import (
	"go-mux-gorm-crud/internal/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Place struct {
	gorm.Model
	Name string `json:"name"`
	Star int    `json:"star"`
}

func InitDB() {
	config.ConnectDatabase()
	db = config.GetDB()
	db.AutoMigrate(&Place{})
}

func GetAllPlaces() ([]Place, error) {
	var places []Place
	err := db.Find(&places).Error
	if err != nil {
		return nil, err
	}
	return places, nil
}

func GetPlace(id int) (Place, error) {
	var place Place
	err := db.First(&place, id).Error
	if err != nil {
		return place, err
	}
	return place, nil
}

func CreatePlace(place Place) (Place, error) {
	err := db.Create(&place).Error
	if err != nil {
		return place, err
	}
	return place, nil
}

func UpdatePlace(id int, place Place) (Place, error) {
	err := db.Model(&place).Where("id = ?", id).Updates(place).Error
	if err != nil {
		return place, err
	}
	return place, nil
}

func DeletePlace(id int) error {
	err := db.Where("id = ?", id).Delete(&Place{}).Error
	if err != nil {
		return err
	}
	return nil
}
