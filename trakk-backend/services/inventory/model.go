package inventory
import (
	"time"
)

type Inventory struct{	
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Quantity    int       `json:"quantity" bson:"quantity"`
	Price       float64   `json:"price" bson:"price"`
	Category    string    `json:"category" bson:"category"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}