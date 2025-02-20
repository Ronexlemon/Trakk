package api

import (
	"log"
	"net/http"
	"trakk/db"
	"github.com/gorilla/mux"
)

var addr = ":9090"

func Server() {
	
	router := mux.NewRouter()

	
	db.CreateClient()

	
	InitializeRoutes(router)

	
	log.Println("Server is running on port", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("ListenAndServe Error:", err)
	}
}
