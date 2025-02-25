package inventory

import (
	"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Inventory struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId      bson.ObjectID `json:"user_id" bson:"user_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Quantity    int           `json:"quantity" bson:"quantity"`
	Price       float64       `json:"price" bson:"price"`
	Category    string        `json:"category" bson:"category"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}
