package auth


type User struct{
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