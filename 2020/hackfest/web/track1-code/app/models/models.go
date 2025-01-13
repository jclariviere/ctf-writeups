package models

import (
	"gorm.io/gorm"
)

// Defining models
type User struct {
	gorm.Model
	IP                   *uint32
	Guid                 *string
	Name                 *string
	Surname              *string
	GalleyRegistrationID *string
	CountryOfOrigin      *string
	MainMaterial         *string
	YearOfFabrication    *uint16
	WeightOfMerchandise  *uint8
	CrewCount            *uint8
	NumberOfPaddles      *uint8
	NumberOfSails        *uint8
	Age                  *uint8
	Active               *uint8
	Admin                *uint8
}

// Migration function
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
