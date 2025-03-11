package inventory

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"trakk/middleware"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type  InventoryRoutes struct{
	repository *Repository
}

var userContext  =middleware.UserContextKey

func CreateRoutes(repo *Repository)*InventoryRoutes{
	return &InventoryRoutes{repository: repo}
}

func (i *InventoryRoutes) CreateInventory(w http.ResponseWriter, r *http.Request){
	
	
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
	
	result,err := i.repository.Create(&inventory,ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"Created":result.Name})

}


func (i *InventoryRoutes) UserInventories(w http.ResponseWriter , r *http.Request){
	claims,ok := r.Context().Value(userContext).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Failed to retrieve user info from context", http.StatusUnauthorized)
		return
		}
		id := claims["id"].(string)
		
		ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
		defer cancel()
		inventories,err:=i.repository.GetAll(id,ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
			}
			json.NewEncoder(w).Encode(map[string][]Inventory{"Inventories":inventories})
	
}


func (i *InventoryRoutes) UpdateInventory(w http.ResponseWriter, r *http.Request){
	var inventory Inventory
	claims,ok := r.Context().Value(userContext).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Failed to retrieve user info from context", http.StatusUnauthorized)
		return
		}
		
		vars := mux.Vars(r) 
inventory_Id := vars["id"]
		
	if inventory_Id == "" {
		http.Error(w, "Missing inventory ID", http.StatusBadRequest)
		return
	}
		ctx,cancel := context.WithTimeout(context.Background(),time.Second *10)
		defer cancel()
		inventory.UpdatedAt = time.Now()
		err:= json.NewDecoder(r.Body).Decode(&inventory)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
			}

		id,err :=  bson.ObjectIDFromHex(claims["id"].(string))
		if err != nil {
			http.Error(w,"Invalid user ID Format", http.StatusBadRequest)
			return
		}
		inventoryId, err := bson.ObjectIDFromHex(inventory_Id)
		if err != nil {
			http.Error(w, "Invalid inventory ID format", http.StatusBadRequest)
			return
		}
		message ,err:=i.repository.Update(id,inventoryId,&inventory,ctx)
		if err !=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message":message})
}




func (i *InventoryRoutes) Deletenventory(w http.ResponseWriter, r *http.Request){
	var inventory Inventory
	claims,ok := r.Context().Value(userContext).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Failed to retrieve user info from context", http.StatusUnauthorized)
		return
		}
		
		vars := mux.Vars(r) 
inventory_Id := vars["id"]
		
	if inventory_Id == "" {
		http.Error(w, "Missing inventory ID", http.StatusBadRequest)
		return
	}
		ctx,cancel := context.WithTimeout(context.Background(),time.Second *10)
		defer cancel()
		inventory.UpdatedAt = time.Now()
		err:= json.NewDecoder(r.Body).Decode(&inventory)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
			}

		id,err :=  bson.ObjectIDFromHex(claims["id"].(string))
		if err != nil {
			http.Error(w,"Invalid user ID Format", http.StatusBadRequest)
			return
		}
		inventoryId, err := bson.ObjectIDFromHex(inventory_Id)
		if err != nil {
			http.Error(w, "Invalid inventory ID format", http.StatusBadRequest)
			return
		}
		message ,err:=i.repository.Delete(id,inventoryId,ctx)
		if err !=nil{
			http.Error(w,"Failed to Delete", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message":message})
}


func (i *InventoryRoutes) InventoriesPerPeriod(w http.ResponseWriter, r *http.Request){
	claims,ok := r.Context().Value(userContext).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Failed to retrieve user info from context", http.StatusUnauthorized)
		return
		}
	ctx, cancel:= context.WithTimeout(context.Background(),5 *time.Second)
	defer cancel()
	id,err :=  bson.ObjectIDFromHex(claims["id"].(string))
		if err != nil {
			http.Error(w,"Invalid user ID Format", http.StatusBadRequest)
			return
		}

	inventories,err:=i.repository.InventoryPerPeriod(id,"monthly",ctx)
	if err !=nil{
		http.Error(w,"Unable to fetch inventorires",http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string][]Inventory{"inventories":inventories})
}