package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rendranarendra/todolist/models"
	b "github.com/rendranarendra/todolist/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = models.GetDB().Collection("users")

// CreateUser controller
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		b.Respond(w, b.Message(false, "Invalid request"))
		return
	}

	resp := user.Create()
	b.Respond(w, resp)
}

// Authenticate controller
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		b.Respond(w, b.Message(false, "Invalid Request"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	b.Respond(w, resp)
}

// CurrentUser controller
var CurrentUser = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(primitive.ObjectID)
	data := models.Current((user))

	resp := b.Message(true, "success")
	resp["data"] = data
	b.Respond(w, resp)
}
