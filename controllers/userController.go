package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rendranarendra/todolist/models"
	b "github.com/rendranarendra/todolist/utils"
)

var collection = models.GetDB().Collection("users")

// CreateUser function
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
