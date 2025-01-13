package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"unsafe"

	"gorm.io/gorm"
	"hackerman.ca/me/models"
	"hackerman.ca/me/utils"
)

func CreateUserEntry(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	guid, guidErr := utils.ValidateSessionID(r, w)

	if guidErr != nil {
		return
	}

	stringRegex := regexp.MustCompile(`^[\w ,.'-0-9]{1,255}$`)

	r.ParseForm()

	var name string = r.Form.Get("name")
	var surname string = r.Form.Get("surname")
	var galleyRegistrationID string = r.Form.Get("galley_registration_id")
	var countryOfOrigin string = r.Form.Get("country_of_origin")
	var mainMaterial string = r.Form.Get("main_material")

	yearOfFabrication, yearOfFabricationErr := strconv.Atoi(r.Form.Get("year_of_fabrication"))

	weightOfMerchandise, weightOfMerchandiseErr := strconv.Atoi(r.Form.Get("weight_of_merchandise"))
	crewCount, crewCountErr := strconv.Atoi(r.Form.Get("crew_count"))
	numberOfPaddles, numberOfPaddlesErr := strconv.Atoi(r.Form.Get("number_of_paddles"))
	numberOfSails, numberOfSailsErr := strconv.Atoi(r.Form.Get("number_of_sails"))
	age, ageErr := strconv.Atoi(r.Form.Get("age"))

	ip, ipErr := utils.GetIP(r)

	if !stringRegex.MatchString(name) || !stringRegex.MatchString(surname) || !stringRegex.MatchString(galleyRegistrationID) || !stringRegex.MatchString(countryOfOrigin) || !stringRegex.MatchString(mainMaterial) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid data")
		return
	}

	if yearOfFabricationErr != nil || weightOfMerchandiseErr != nil || crewCountErr != nil || numberOfPaddlesErr != nil || numberOfSailsErr != nil || ageErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid data")
		return
	}

	if ipErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, fmt.Sprintf("%s", ipErr))
		return
	}

	yearOfFabricationUint16 := uint16(yearOfFabrication)
	weightOfMerchandiseUint8 := uint8(weightOfMerchandise)
	crewCountUint8 := uint8(crewCount)
	numberOfPaddlesUint8 := uint8(numberOfPaddles)
	numberOfSailsUint8 := uint8(numberOfSails)
	ageUint8 := uint8(age)

	encodedIP := utils.EncodeIP(ip)

	active := uint8(1)
	admin := uint8(0)

	db.Create(&models.User{IP: &encodedIP, Guid: &guid, Name: &name, Surname: &surname, GalleyRegistrationID: &galleyRegistrationID, CountryOfOrigin: &countryOfOrigin, MainMaterial: &mainMaterial, YearOfFabrication: &yearOfFabricationUint16, WeightOfMerchandise: &weightOfMerchandiseUint8, CrewCount: &crewCountUint8, NumberOfPaddles: &numberOfPaddlesUint8, NumberOfSails: &numberOfSailsUint8, Age: &ageUint8, Active: &active, Admin: &admin})
}

func ReadUserEntry(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	guid, guidErr := utils.ValidateSessionID(r, w)

	if guidErr != nil {
		return
	}

	var user models.User

	db.First(&user, "guid = ?", guid)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	webUser := utils.UserWebRepresentation{IP: utils.DecodeIP(*user.IP), Name: *user.Name, Surname: *user.Surname, GalleyRegistrationID: *user.GalleyRegistrationID, CountryOfOrigin: *user.CountryOfOrigin, MainMaterial: *user.MainMaterial, YearOfFabrication: *user.YearOfFabrication, WeightOfMerchandise: *user.WeightOfMerchandise, CrewCount: *user.CrewCount, NumberOfPaddles: *user.NumberOfPaddles, NumberOfSails: *user.NumberOfSails, Age: *user.Age, Active: *user.Active, Admin: *user.Admin}

	response, responseErr := json.Marshal(webUser)

	if responseErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occured")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)

	return
}

func DeleteUserEntry(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	guid, guidErr := utils.ValidateSessionID(r, w)

	if guidErr != nil {
		return
	}

	var user models.User

	db.First(&user, "guid = ?", guid)
	db.Model(&user).Update("Active", 0)

	w.WriteHeader(http.StatusOK)
	return

}

func EditUserEntry(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	guid, guidErr := utils.ValidateSessionID(r, w)

	if guidErr != nil {
		return
	}

	var user models.User

	db.First(&user, "guid = ?", guid)

	stringRegex := regexp.MustCompile(`^[\w ,.'-]{1,255}$`)

	r.ParseForm()

	var name string = r.Form.Get("name")
	var surname string = r.Form.Get("surname")
	var galleyRegistrationID string = r.Form.Get("galley_registration_id")
	var countryOfOrigin string = r.Form.Get("country_of_origin")
	var mainMaterial string = r.Form.Get("main_material")

	yearOfFabrication, yearOfFabricationErr := strconv.Atoi(r.Form.Get("year_of_fabrication"))

	weightOfMerchandise, weightOfMerchandiseErr := strconv.Atoi(r.Form.Get("weight_of_merchandise"))
	crewCount, crewCountErr := strconv.Atoi(r.Form.Get("crew_count"))
	numberOfPaddles, numberOfPaddlesErr := strconv.Atoi(r.Form.Get("number_of_paddles"))
	numberOfSails, numberOfSailsErr := strconv.Atoi(r.Form.Get("number_of_sails"))
	age, ageErr := strconv.Atoi(r.Form.Get("age"))

	ip, ipErr := utils.GetIP(r)

	if !stringRegex.MatchString(name) || !stringRegex.MatchString(surname) || !stringRegex.MatchString(galleyRegistrationID) || !stringRegex.MatchString(countryOfOrigin) || !stringRegex.MatchString(mainMaterial) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid data")
		return
	}

	if yearOfFabricationErr != nil || weightOfMerchandiseErr != nil || crewCountErr != nil || numberOfPaddlesErr != nil || numberOfSailsErr != nil || ageErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid data")
		return
	}

	if ipErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, fmt.Sprintf("%s", ipErr))
		return
	}

	yearOfFabricationUint16 := uint16(yearOfFabrication)

	var userParamsArray [7]uint8

	userParamsArray[0] = uint8(crewCount)
	userParamsArray[1] = uint8(numberOfPaddles)
	userParamsArray[2] = uint8(numberOfSails)
	userParamsArray[3] = uint8(age)
	userParamsArray[5] = *user.Active
	userParamsArray[6] = *user.Admin

	var weightOfMerchandiseCasted interface{}

	if weightOfMerchandise < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid weight")
		return
	}
	if weightOfMerchandise <= 255 {
		weightOfMerchandiseCasted = uint8(weightOfMerchandise)
		weightOfMerchandiseUint8, _ := weightOfMerchandiseCasted.(uint8)
		userParamsArray[4] = weightOfMerchandiseUint8
	} else if weightOfMerchandise <= 65535 {
		weightOfMerchandiseCasted = uint16(weightOfMerchandise)
		weightOfMerchandiseUint16, _ := weightOfMerchandiseCasted.(uint16)
		*(*uint16)(unsafe.Pointer(&userParamsArray[4])) = weightOfMerchandiseUint16
	} else if weightOfMerchandise <= 4294967295 {
		weightOfMerchandiseCasted = uint32(weightOfMerchandise)
		weightOfMerchandiseUint32, _ := weightOfMerchandiseCasted.(uint32)
		*(*uint32)(unsafe.Pointer(&userParamsArray[4])) = weightOfMerchandiseUint32
	} else {
		weightOfMerchandiseCasted = uint64(weightOfMerchandise)
		weightOfMerchandiseUint64, _ := weightOfMerchandiseCasted.(uint64)
		*(*uint64)(unsafe.Pointer(&userParamsArray[4])) = weightOfMerchandiseUint64
	}

	encodedIP := utils.EncodeIP(ip)

	db.Model(&user).Updates(models.User{IP: &encodedIP, Name: &name, Surname: &surname, GalleyRegistrationID: &galleyRegistrationID, CountryOfOrigin: &countryOfOrigin, MainMaterial: &mainMaterial, YearOfFabrication: &yearOfFabricationUint16, WeightOfMerchandise: &userParamsArray[4], CrewCount: &userParamsArray[0], NumberOfPaddles: &userParamsArray[1], NumberOfSails: &userParamsArray[2], Age: &userParamsArray[3], Active: &userParamsArray[5], Admin: &userParamsArray[6]})

}
