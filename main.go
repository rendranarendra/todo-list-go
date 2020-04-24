package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rendranarendra/todolist/app"
	"github.com/rendranarendra/todolist/controllers"
	"github.com/gorilla/mux"
	"github.com/maple-ai/syrup"
)

func main() {
	router := syrup.New(mux.NewRouter())

	//User router
	router.Post("/api/user/create", controllers.CreateUser)
	router.Post("/api/user/login", controllers.Authenticate)
	router.Get("/api/user/me", controllers.CurrentUser)

	//Post router
	router.Post("/api/post", controllers.CreatePost)

	//Middleware
	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	fmt.Println("Listening on port ", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
