package auth

import (
	"context"
	"fmt"
	"strings"
	"trakk/db"

	//supa "github.com/nedpals/supabase-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


type UserRepository struct {
	//dbclient *supa.Client
	mongoclient *mongo.Client
}

func NewUserRepository() *UserRepository {
	return &UserRepository{mongoclient: db.MongoClient}
}

func (u *UserRepository) Create(user *User,ctx context.Context) (string, error) {
	
	fmt.Println("Inserting user:", user)
	
	result ,err:= u.mongoclient.Database("Trakk").Collection("users").InsertOne(ctx,user)
	if err != nil {
		return "Failed to add user", err
		}
	

	fmt.Println("User inserted successfully:", result)
	return user.Email, nil
}

func (u *UserRepository) checkUser(email string, password string,ctx context.Context) (User, error) {
	var user User
	

	
	email = strings.TrimSpace(email)

	

	
	
	err := u.mongoclient.Database("Trakk").Collection("users").FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&user)


	
	if err != nil {
		fmt.Println("Error querying user:", err)
		return User{}, err
	}

	

	// Verify the password  and decrypt
	if user.Password != password { // Ensure this matches your password storage method (e.g., hashed passwords)
		fmt.Println("Invalid password for email:", email)
		return User{}, fmt.Errorf("invalid password")
	}

	// Return the user details
	return user, nil
}