package routes

import (
	"fmt"
	"go-api/users"

	"github.com/gorilla/mux"
)


func InitRouter() *mux.Router{
  fmt.Println("initializing router")
  r := mux.NewRouter()
  r.Use(mux.CORSMethodMiddleware(r))

  //define routes 
  r.HandleFunc("/users/signup",users.CreateUserHandler).Methods("POST")

  return r
}
