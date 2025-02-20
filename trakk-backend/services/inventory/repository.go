package inventory

import (
	"trakk/db"

	supa "github.com/nedpals/supabase-go"
)


type Repository struct {
	dbClient *supa.Client
}


func NewRepository()*Repository{
	return &Repository{dbClient: db.Supabase}
}

func (r *Repository) Create(inventory *Inventory)(Inventory,error){
	//logic to create inventory
	
	return *inventory,nil
}

func (r *Repository) Delete(inventory_id string)(string,error){
	//logic to delete inventory
	return "inventory deleted",nil
}
func (r *Repository) Update(inventory_id string)(string, error){
	//logic to update inventory
	return "inventory updated",nil
}

func (r *Repository) GetAll(user_id string)(*Inventory,error){
	//logic to get all inventory
	return &Inventory{},nil
}