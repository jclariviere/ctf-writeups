package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hackerman.ca/me/handlers"
	"hackerman.ca/me/models"
)

// Router wrapper
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Function to run the app
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// DB initialization
func (a *App) Initialize() {

	var db *gorm.DB
	var err error

	if os.Getenv("POSTGRES_HOST") != "" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Toronto", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = models.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// HTTP methods wrappers
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Route declarations
func (a *App) setRouters() {
	a.Get("/api/get-id", a.SetSessionCookie)
	a.Post("/api/fill-a-38", a.CreateUserEntry)
	a.Get("/api/get-a-38", a.ReadUserEntry)
	a.Delete("/api/delete-a-38", a.DeleteUserEntry)
	a.Post("/api/edit-a-38", a.EditUserEntry)
	a.Get("/api/get-yellow", a.AdminFlag)
	a.Get("/api/get-pink", a.MFAFlag)
	a.Get("/api/get-brown", a.IPFlag)
}

// API functions wrappers
func (a *App) SetSessionCookie(w http.ResponseWriter, r *http.Request) {
	handlers.SetSessionCookie(a.DB, w, r)
}
func (a *App) CreateUserEntry(w http.ResponseWriter, r *http.Request) {
	handlers.CreateUserEntry(a.DB, w, r)
}
func (a *App) ReadUserEntry(w http.ResponseWriter, r *http.Request) {
	handlers.ReadUserEntry(a.DB, w, r)
}
func (a *App) DeleteUserEntry(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteUserEntry(a.DB, w, r)
}
func (a *App) EditUserEntry(w http.ResponseWriter, r *http.Request) {
	handlers.EditUserEntry(a.DB, w, r)
}
func (a *App) AdminFlag(w http.ResponseWriter, r *http.Request) {
	handlers.AdminFlag(a.DB, w, r)
}
func (a *App) MFAFlag(w http.ResponseWriter, r *http.Request) {
	handlers.MFAFlag(a.DB, w, r)
}
func (a *App) IPFlag(w http.ResponseWriter, r *http.Request) {
	handlers.IPFlag(a.DB, w, r)
}
