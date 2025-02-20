package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserRoutes struct{
	userRepo *UserRepository
}

func  InitializeUserRoutes(userrepo *UserRepository)*UserRoutes{
	return &UserRoutes{userRepo: userrepo}

}

func (u *UserRoutes) CreateUser(w http.ResponseWriter, r *http.Request){
	//logic to create user
	var user  User
	err:= json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return}
	message,err:=u.userRepo.Create(&user)
	fmt.Println("user created message",message,err)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return}
		json.NewEncoder(w).Encode(message)
}