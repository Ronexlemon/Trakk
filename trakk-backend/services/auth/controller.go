package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"trakk/pkg/auth"
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
		ctx,cancel := context.WithTimeout(context.Background(),5 *time.Second)
		defer cancel()
	message,err:=u.userRepo.Create(&user,ctx)
	fmt.Println("user created message",message,err)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return}
		json.NewEncoder(w).Encode(message)
}


func (u *UserRoutes) Login( w http.ResponseWriter, r *http.Request){
	//logic to login
	var user User
	err:=json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("User deatails",user)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return}
		ctx,cancel := context.WithTimeout(context.Background(),5 *time.Second)
		defer cancel()
	userr,err:=u.userRepo.login(user.Email,user.Password,ctx)
	fmt.Println("User two details",userr)
	if err !=nil{
		http.Error(w,err.Error(),http.StatusUnauthorized)
		return
	}
	//jwt create token
	token,err := auth.CreateToken(userr.Username,user.Email,userr.Phone)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
		}
		json.NewEncoder(w).Encode(token)
		

	

}