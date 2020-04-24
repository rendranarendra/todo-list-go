package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rendranarendra/todolist/models"
	u "github.com/rendranarendra/todolist/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePost function
var CreatePost = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(primitive.ObjectID)
	post := &models.Post{}

	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while processing body"))
		return
	}

	post.Owner = userID
	resp := post.Create()
	u.Respond(w, resp)
}
