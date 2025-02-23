package api

import (
	"trakk/services/auth"
	"trakk/services/inventory"

	"github.com/gorilla/mux"
)

func InitializeRoutes(r *mux.Router){	
  i:=inventory.NewRepository()
  inventoryRoutes :=inventory.CreateRoutes(i)
  u:=auth.NewUserRepository()
  authRoutes := auth.InitializeUserRoutes(u)
  r.HandleFunc("/api/user/create",authRoutes.CreateUser).Methods("POST")
  r.HandleFunc("/api/user/login",authRoutes.Login).Methods("POST")
  r.HandleFunc("/api/inventory,create",inventoryRoutes.CreateInventory).Methods("POST")
  //r.HandleFunc("/api/user/login",authRoutes.LoginUser).Methods("POST")


}