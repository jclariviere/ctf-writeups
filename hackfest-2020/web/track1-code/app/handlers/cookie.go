package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SetSessionCookie(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	_, e := r.Cookie("session_id")

	if e == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	cookie := http.Cookie{
		Name:  "session_id",
		Value: uuid.New().String(),
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	return
}
