package inventory

import (
	"context"
	"fmt"
	"trakk/db"

	//supa "github.com/nedpals/supabase-go"
	//"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (r *Repository) GetAll(user_id string,ctx context.Context)([]Inventory,error){
	fmt.Println("user_id",user_id)
	id,err :=bson.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, err
		}
	cursor,err := r.dbClient.Database("Trakk").Collection("inventories").Find(ctx,bson.M{"user_id":id})
	fmt.Println("cursor",cursor)
	if err != nil{
		return nil,err
		}
		var inventories []Inventory
		if err = cursor.All(ctx,&inventories);err != nil{
			return nil,err
			}
			return inventories,nil
	
}