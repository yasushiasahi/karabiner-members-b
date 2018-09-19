package main

import (
	"log"
	"net/http"

	"github.com/yasushiasahi/karabiner-members/api/db"
	"github.com/yasushiasahi/karabiner-members/api/handler"
)

func main() {

	master, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	uh := handler.NewUserHandler(master)
	ush := handler.NewUsersHandler(master)
	msh := handler.MembersHandler{}

	mux := http.NewServeMux()
	mux.Handle("/api/user", uh)
	mux.Handle("/api/users", ush)
	mux.Handle("/api/members", msh)

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	log.Printf("server started")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
