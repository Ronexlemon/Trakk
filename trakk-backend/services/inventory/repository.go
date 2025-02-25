package inventory

import (
	"context"
	"trakk/db"

	//supa "github.com/nedpals/supabase-go"
	//"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


type Repository struct {
	//dbClient *supa.Client
	dbClient  *mongo.Client
}


func NewRepository()*Repository{
	return &Repository{dbClient: db.MongoClient}
}

func (r *Repository) Create(inventory *Inventory,ctx context.Context)(Inventory,error){
	//logic to create inventory
	_,err := r.dbClient.Database("Trakk").Collection("inventories").InsertOne(ctx,inventory)
	if err !=nil{
		return Inventory{},err
	}

	
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