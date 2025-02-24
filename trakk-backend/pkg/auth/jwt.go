package auth

import (
	"log"
	"time"
	"trakk/config"

	"github.com/golang-jwt/jwt/v5"
)



func CreateToken(username string,email string , phone string)(string,error){
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"username":username,
		"email":email,
		"phone":phone,
		"exp":jwt.NewNumericDate(time.Now().Add(time.Hour*72)).Unix(),
		"iat": time.Now().Unix(),

	})
	secret, err := config.LoadJWTsecret()
	if err != nil {
		log.Fatal(err)
		return "",err
		}
		tokenString, err := token.SignedString(secret)
		if err != nil {
			return "",err
			}
			return tokenString,nil

}

func VerifyJwt