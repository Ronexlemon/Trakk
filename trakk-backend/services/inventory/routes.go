package inventory

import (
	"encoding/json"
	"net/http"
)

type  InventoryRoutes struct{
	repository *Repository
}

func CreateRoutes(repo *Repository)*InventoryRoutes{
	return &InventoryRoutes{repository: repo}
}

func (i *InventoryRoutes) CreateInventory(w http.ResponseWriter, r *http.Request){
	var inventory Inventory

	err:= json.NewDecoder(r.Body).Decode(&inventory)
	if err !=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//TODO  pass the data to the db

}
