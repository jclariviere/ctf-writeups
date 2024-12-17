package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type UserWebRepresentation struct {
	IP                   string
	Name                 string
	Surname              string
	GalleyRegistrationID string
	CountryOfOrigin      string
	MainMaterial         string
	YearOfFabrication    uint16
	WeightOfMerchandise  uint8
	CrewCount            uint8
	NumberOfPaddles      uint8
	NumberOfSails        uint8
	Age                  uint8
	Active               uint8
	Admin                uint8
}

func ValidateSessionID(r *http.Request, w http.ResponseWriter) (string, error) {
	guidRegex := regexp.MustCompile(`^[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$`)

	guid, guidErr := r.Cookie("session_id")

	if guidErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, fmt.Sprintf("%s", "No session id"))
		return "", errors.New("No session id")
	}

	var err error

	if !guidRegex.MatchString(guid.Value) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, fmt.Sprintf("%s", "Invalid session id"))
		err = errors.New("Invalid session id")
	}

	return guid.Value, err
}

func ValidateMFA(mfaInput string, mfaType *string, mfaOK *bool) {

	*mfaOK = true

	var mfaCode [9]int

	for i := 0; i < 9; i++ {
		mfaCode[i] = rand.Intn(9)
	}

	if *mfaType == "compact" {
		for i, digit := range mfaInput {
			digitInt, _ := strconv.Atoi(string(digit))
			if mfaCode[i] != digitInt {
				*mfaOK = false
			}
		}
	} else if *mfaType == "space" {
		inputWIthoutSpace := strings.Replace(mfaInput, " ", "", -1)
		for i, digit := range inputWIthoutSpace {
			digitInt, _ := strconv.Atoi(string(digit))
			if mfaCode[i] != digitInt {
				*mfaOK = false
			}
		}
	} else if *mfaType == "dash" {
		inputWIthoutDash := strings.Replace(mfaInput, "-", "", -1)
		for i, digit := range inputWIthoutDash {
			digitInt, _ := strconv.Atoi(string(digit))
			if mfaCode[i] != digitInt {
				*mfaOK = false
			}
		}
	} else {
		*mfaOK = false
	}

}

func MFAVerdict(w http.ResponseWriter, mfaOK *bool) {

	flag3 := os.Getenv("FLAG3")

	if *mfaOK {
		recover()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, flag3)
		return
	}

	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintf(w, "Sorry, the MFA code is incorrect")
	return
}

func GetIP(r *http.Request) (string, error) {
	forwardedSlice := strings.Split(r.Header.Get("X-Forwarded-For"), ", ")
	forwarded := forwardedSlice[0]
	ipRegex := regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)

	quadZeroRegex := regexp.MustCompile(`0+\.0+\.0+\.0+`)

	if quadZeroRegex.MatchString(forwarded) {
		return r.RemoteAddr, errors.New("Sorry but this is very unsafe and therefor cannot be allowed")
	}

	if ipRegex.MatchString(forwarded) {
		return forwarded, nil
	}
	return r.RemoteAddr, nil
}

func EncodeIP(rawIP string) (decimalIP uint32) {

	ipSlice := strings.Split(rawIP, ".")

	pos := 24

	for _, octet := range ipSlice {
		castedOctet, _ := strconv.Atoi(octet)
		decimalIP += uint32(castedOctet) << pos
		pos -= 8
	}
	return
}

func DecodeIP(decimalIP uint32) (rawIP string) {

	pos := 24

	var ipArray []string

	for i := 0; i < 4; i++ {
		ipArray = append(ipArray, strconv.Itoa(int(uint8(decimalIP>>pos))))
		pos -= 8
	}

	rawIP = fmt.Sprintf(strings.Join(ipArray, "."))
	return
}
