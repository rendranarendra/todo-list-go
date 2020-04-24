package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rendranarendra/todolist/controllers"
	"github.com/gorilla/mux"
	"github.com/maple-ai/syrup"
)

func main() {
	router := syrup.New(mux.NewRouter())

	//User router
	router.Post("/api/user", controllers.CreateUser)

	//Post router
	router.Post("/api/post", controllers.CreatePost)

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
