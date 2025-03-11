package inventory

import (
	"context"
	"fmt"
	"time"
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


func (r *Repository) InventoryPerPeriod(user_id bson.ObjectID,year int, month int, day *int, periodType string, ctx context.Context)([]Inventory,error){
	if month < 1 || month > 12 {
		return nil, fmt.Errorf("invalid month value, it should be between 1 and 12")
	}
	var matchcondition bson.M
	var startDate, endDate time.Time
	

	switch periodType{
	case "day":
		if day == nil || *day < 1 || *day > 31 {
			return nil, fmt.Errorf("invalid day value")
		}
		// Set start and end to the same day
		startDate = time.Date(year, time.Month(month), *day, 0, 0, 0, 0, time.UTC)
		endDate = startDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	case "month":
		startDate = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		endDate = startDate.AddDate(0, 1, -1).Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	case "year":
		startDate = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		endDate = time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	case "6months":
		endDate = time.Now()
		startDate = endDate.AddDate(0, -6, 0)
	default:
		return nil, fmt.Errorf("invalid period type. Use 'monthly' or 'weekly'")
	}
	matchcondition = bson.M{
		"user_id": user_id,
		"created_at":bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}
	pipeline :=[]bson.M{
		{"$match":matchcondition},
		{
			"$project": bson.M{
				"_id":         1,          // include the _id field
				"user_id":     1,          // Include user_id
				"name":        1,          // Include name
				"description": 1,          // Include description
				"quantity":    1,          // Include quantity
				"price":       1,          // Include price
				"category":    1,          // Include category
				"created_at":  1,          // Include created_at
				"updated_at":  1,          // Include updated_at
			},},
	}
	cursor,err :=r.dbClient.Database("Trakk").Collection("inventories").Aggregate(ctx,pipeline)
	
	if err !=nil{
		return nil, err
	}
	defer cursor.Close(ctx)
	var inventories []Inventory
	for cursor.Next(ctx){
		var inventory Inventory
		if err := cursor.Decode(&inventory); err !=nil{
			return nil,err
		}
		inventories = append(inventories, inventory)
	}
	if err :=cursor.Err();err !=nil{
		return nil,err
	}
	return inventories, nil

}