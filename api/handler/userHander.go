package handler

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// UserHandler handles api for users
type UserHandler struct {
	master *gorm.DB
}

// NewUserHandler returns UserHandler
func NewUserHandler(master *gorm.DB) http.Handler {
	return &UserHandler{master: master}
}

// ServeHTTP ...
func (h UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "POST":
		h.create(w, r)
		return
	default:
		http.NotFound(w, r)
	}
}

// GetUser ...
func (h UserHandler) create(w http.ResponseWriter, r *http.Request) {
	u, err := decodeBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	h.master.Create(&u)
	w.WriteHeader(http.StatusOK)
}
