package models

import (
	"context"

	"github.com/Kamva/mgm"
	"github.com/dgrijalva/jwt-go"
	u "github.com/rendranarendra/todolist/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Token            string `json:"token"`
}

// Token JWT model
type Token struct {
	UserID primitive.ObjectID
	jwt.StandardClaims
}

//Create new user
func (user *User) Create() map[string]interface{} {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(encryptedPassword)

	collection := GetDB().Collection("users")
	collection.InsertOne(context.TODO(), user)

	response := u.Message(true, "User has been created")
	response["user"] = user
	return response
}
