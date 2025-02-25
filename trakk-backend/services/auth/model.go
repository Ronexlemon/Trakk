package auth

import (
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)


type User struct{
	ID       bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"`
}

type LoginUser struct {
    Email    string `json:"email" bson:"email"`
    Password string `json:"password" bson:"password"`
	Username string `json:"username" bson:"username"`
}