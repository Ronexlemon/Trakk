package api

import (
	"net/http"
	"trakk/middleware"
	"trakk/services/auth"
	"trakk/services/inventory"

	"github.com/gorilla/mux"
)

func InitializeRoutes(r *mux.Router){	
  i:=inventory.NewRepository()
  inventoryRoutes :=inventory.CreateRoutes(i)
  u:=auth.NewUserRepository()
  authRoutes := auth.InitializeUserRoutes(u)
  r.HandleFunc("/api/user/signup",authRoutes.CreateUser).Methods("POST")
  r.HandleFunc("/api/user/login",authRoutes.Login).Methods("POST")
  r.Handle("/api/inventory/create",middleware.JwtAuthMiddleware(http.HandlerFunc(inventoryRoutes.CreateInventory))).Methods("POST")
  r.Handle("/api/inventory/user/inventories",middleware.JwtAuthMiddleware(http.HandlerFunc(inventoryRoutes.UserInventories))).Methods("GET")

  //r.HandleFunc("/api/user/login",authRoutes.LoginUser).Methods("POST")


}