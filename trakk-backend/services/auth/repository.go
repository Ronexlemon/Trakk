package auth

import (
	"context"
	"fmt"
	"strings"
	"trakk/db"
	"trakk/pkg/encrption"

	//supa "github.com/nedpals/supabase-go"
	
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


type UserRepository struct {
	//dbclient *supa.Client
	mongoclient *mongo.Client
}

func NewUserRepository() *UserRepository {
	db.CreateUserIndex(db.MongoClient) //create unique fields
	return &UserRepository{mongoclient: db.MongoClient}
}

func (u *UserRepository) Create(user *User, ctx context.Context) (string, error) {
	
	hashPassword, err := encrption.HashPassword(user.Password)
	if err != nil {
		return "failed to hash password", err
	}
	user.Password = hashPassword

	
	use_r, err := u.checkUser(user.Email, ctx)
	if err == nil {
		
		return fmt.Sprintf("User Already Exists: %s", use_r.Email), nil
	}

	
	result, err := u.mongoclient.Database("Trakk").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return "Failed to add user", err
	}

	
	fmt.Println("User inserted successfully:", result)
	return user.Email, nil
}

func (u *UserRepository) checkUser(email string,ctx context.Context) (User, error) {
	var user User
	


	
	email = strings.TrimSpace(email)

	

	
	
	err := u.mongoclient.Database("Trakk").Collection("users").FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		
		return User{}, fmt.Errorf("user not found")
	}
	
	if err != nil {
		fmt.Println("Error querying user:", err)
		return User{}, err
	}
	
	

	
	user.Password =""

	// Return the user details
	return user, nil
}

func (u *UserRepository) login(email string, password string,ctx context.Context) (User, error) {
	var user User
	

	
	email = strings.TrimSpace(email)

	

	
	
	err := u.mongoclient.Database("Trakk").Collection("users").FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		
		return User{}, fmt.Errorf("user not found")
	}
	
	if err != nil {
		fmt.Println("Error querying user:", err)
		return User{}, err
	}
	
	

	// Verify the password  and decrypt
	
	if !encrption.VerifyPassword(user.Password,password) { 
		fmt.Println("Invalid password for email:", email)
		return User{}, fmt.Errorf("invalid password")
	}
	user.Password =""
	

	// Return the user details
	return user, nil
}