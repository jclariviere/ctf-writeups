package handlers

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"gorm.io/gorm"
	"hackerman.ca/me/models"
	"hackerman.ca/me/utils"
)

func AdminFlag(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	guid, guidErr := utils.ValidateSessionID(r, w)

	if guidErr != nil {
		return
	}

	flag2 := os.Getenv("FLAG2")
	var user models.User

	db.First(&user, "guid = ?", guid)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if *user.Admin == 1 && *user.Active == 1 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, flag2)
		return
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "You need to be admin and to be active")
		return
	}

}

func MFAFlag(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	mfaOK := false

	defer utils.MFAVerdict(w, &mfaOK)

	var mfaType *string

	formatsArray := [3]string{"compact", "space", "dash"}
	formatsRegex := make(map[string]*regexp.Regexp)

	formatsRegex["main"] = regexp.MustCompile(`^\d{3}[\s-]?\d{3}[\s-]?\d{3}$`)
	formatsRegex["compact"] = regexp.MustCompile(`^\d{9}$`)
	formatsRegex["space"] = regexp.MustCompile(`^\d{3}\s\d{3}\s\d{3}$`)
	formatsRegex["dash"] = regexp.MustCompile(`^\d{3}-\d{3}-\d{3}$`)

	input := r.FormValue("mfa_string")

	if !formatsRegex["main"].MatchString(input) {
		return
	}

	if formatsRegex["compact"].MatchString(input) {
		mfaType = &formatsArray[0]
	} else if formatsRegex["space"].MatchString(input) {
		mfaType = &formatsArray[1]
	} else if formatsRegex["dash"].MatchString(input) {
		mfaType = &formatsArray[2]
	}

	utils.ValidateMFA(input, mfaType, &mfaOK)

}

func IPFlag(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	guid, guidErr := utils.ValidateSessionID(r, w)

	if guidErr != nil {
		return
	}

	flag4 := os.Getenv("FLAG4")
	var user models.User

	db.First(&user, "guid = ?", guid)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if utils.DecodeIP(*user.IP) == "0.0.0.0" || utils.DecodeIP(*user.IP) == "TO BE IMPLEMENTED" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, flag4)
		return
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Sorry, your IP does not match the one stored in the database. ")
		return
	}

}
