package routes

import (
	"go-api/users"

	"github.com/gorilla/mux"
)


func InitRouter() *mux.Router{
  r := mux.NewRouter()
  r.Use(mux.CORSMethodMiddleware(r))

  //define routes 
  r.HandleFunc("/users/signup",users.CreateUserHandler).Methods("POST")

  return r
}
