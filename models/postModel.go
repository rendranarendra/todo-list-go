package models

import (
	"context"

	"github.com/Kamva/mgm"
	u "github.com/rendranarendra/todolist/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post model
type Post struct {
	mgm.DefaultModel `bson:",inline"`
	Owner            primitive.ObjectID `json:"owner"`
	Title            string             `json:"title"`
	DueDate          string             `json:"dueDate"`
	Importance       bool                `json:"importance"`
	Completed        bool               `json:"false"`
}

//Create post
func (post *Post) Create() map[string]interface{} {
	collection := GetDB().Collection("posts")
	collection.InsertOne(context.TODO(), post)

	response := u.Message(true, "Post has been created")
	response["post"] = post
	return response
}
