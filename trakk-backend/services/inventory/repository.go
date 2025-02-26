package inventory

import (
	"context"
	"fmt"
	"trakk/db"

	//supa "github.com/nedpals/supabase-go"
	//"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	//dbClient *supa.Client
	dbClient *mongo.Client
}

func NewRepository() *Repository {
	return &Repository{dbClient: db.MongoClient}
}

func (r *Repository) Create(inventory *Inventory, ctx context.Context) (Inventory, error) {
	//logic to create inventory
	_, err := r.dbClient.Database("Trakk").Collection("inventories").InsertOne(ctx, inventory)
	if err != nil {
		return Inventory{}, err
	}

	return *inventory, nil
}

func (r *Repository) Delete(user_id bson.ObjectID, inventory_id bson.ObjectID, ctx context.Context) (string, error) {
    
    filter := bson.M{"user_id": user_id, "_id": inventory_id}

    
    result, err := r.dbClient.Database("Trakk").Collection("inventories").DeleteOne(ctx, filter)
    if err != nil {
        return "failed to delete", err
    }

    
    if result.DeletedCount == 0 {
        return "No Inventory found", fmt.Errorf("no inventory found")
    }

    
    return "inventory deleted", nil
}
func (r *Repository) Update(user_id bson.ObjectID, inventory_id bson.ObjectID, inventory *Inventory, ctx context.Context) (string, error) {
    invMap, err := bson.Marshal(inventory)
    if err != nil {
        return "failed to marshal inventory", err
    }

    var invData bson.M
    err = bson.Unmarshal(invMap, &invData)
    if err != nil {
        return "failed to unmarshal inventory", err
    }

    delete(invData, "_id")
    delete(invData, "user_id")
	delete(invData, "created_at")

    filter := bson.M{"user_id": user_id, "_id": inventory_id}
    update := bson.M{"$set": invData}

    opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

    var updatedDocument bson.M
    err = r.dbClient.Database("Trakk").Collection("inventories").FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDocument)
    if err != nil {
        return fmt.Sprintf("failed to update: %v", err), err
    }

    return "inventory updated", nil
}

func (r *Repository) GetAll(user_id string, ctx context.Context) ([]Inventory, error) {
	fmt.Println("user_id", user_id)
	id, err := bson.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, err
	}
	cursor, err := r.dbClient.Database("Trakk").Collection("inventories").Find(ctx, bson.M{"user_id": id})
	fmt.Println("cursor", cursor)
	if err != nil {
		return nil, err
	}
	var inventories []Inventory
	if err = cursor.All(ctx, &inventories); err != nil {
		return nil, err
	}
	return inventories, nil

}
