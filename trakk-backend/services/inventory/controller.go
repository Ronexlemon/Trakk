package inventory

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"trakk/middleware"

	"github.com/golang-jwt/jwt/v5"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type  InventoryRoutes struct{
	repository *Repository
}

func CreateRoutes(repo *Repository)*InventoryRoutes{
	return &InventoryRoutes{repository: repo}
}

func (i *InventoryRoutes) CreateInventory(w http.ResponseWriter, r *http.Request){
	userContext :=middleware.UserContextKey
	
	claims, ok := r.Context().Value(userContext).(jwt.MapClaims)
	
	if !ok {
		fmt.Println("Error",ok)
		http.Error(w, "Failed to retrieve user info from context", http.StatusUnauthorized)
		return
	}

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	
	id := claims["id"].(string)
	var inventory Inventory

	err:= json.NewDecoder(r.Body).Decode(&inventory)
	inventory.CreatedAt = time.Now()
	inventory.UpdatedAt = time.Now()
	if err !=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	inventory.UserId,err = bson.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
		}
	//TODO  pass the data to the db
	result,err := i.repository.Create(&inventory,ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"Created":result.Name})

}
