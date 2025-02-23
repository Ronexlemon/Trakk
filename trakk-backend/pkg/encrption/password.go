package encrption

import (
	

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string)(string, error){
	// encrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "", err
		}
		return string(hash), nil

}

func VerifyPassword(passwordHash string,password string)bool{
	err:= bcrypt.CompareHashAndPassword([]byte(passwordHash),[]byte(password))
	if err !=  nil{
		return false
	}
	return true
}
