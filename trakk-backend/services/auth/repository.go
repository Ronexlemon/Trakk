package auth

import (
	"fmt"
	"trakk/db"

	supa "github.com/nedpals/supabase-go"
)
//714888 144479
type UserRepository struct {
	dbclient *supa.Client
}

func NewUserRepository() *UserRepository {
	return &UserRepository{dbclient: db.Supabase}
}

func (u *UserRepository) Create(user *User) (string, error) {
	fmt.Println("Inserting user:", user)
	var result []User
	err := u.dbclient.DB.From("users").Insert(user).Execute(&result)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return "", err
	}

	fmt.Println("User inserted successfully:", result)
	return user.Email, nil
}