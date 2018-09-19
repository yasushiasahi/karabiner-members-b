package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/yasushiasahi/karabiner-members/api/db"
)

// UsersHandler handles api for users
type UsersHandler struct {
	master *gorm.DB
}

// NewUsersHandler returns UserHandler
func NewUsersHandler(master *gorm.DB) http.Handler {
	return &UsersHandler{master: master}
}

// ServeHTTP ...
func (h UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	default:
		http.NotFound(w, r)
	}
}

// GetUsers ...
func (h UsersHandler) get(w http.ResponseWriter, r *http.Request) {
	var us []db.User
	h.master.Order("score desc").Limit(10).Find(&us)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(us)
	if err != nil {
		http.Error(w, "failed to Encode json"+err.Error(), http.StatusInternalServerError)
	}
}
