package config

import (
	"os"

	"github.com/joho/godotenv"
)


func LoadJWTsecret()(string,error){
	err := godotenv.Load()

	if err !=nil{
		return "",err

	}
	jwtsecret := os.Getenv("JWTSECRECT")
	
	return jwtsecret,nil
}