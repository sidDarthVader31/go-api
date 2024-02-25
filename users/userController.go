package users

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, ){
  var nu models.User
  err := json.NewDecoder(r.Body).Decode(&nu)
  fmt.Println("nu:", nu)
  if err != nil{
    fmt.Println("invalid data")
  }
  err1 := createUserService(&nu)
  if err1 != nil{
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("error: %v", err)))
  }
  w.WriteHeader(http.StatusOK)
  jsonResponse, je := json.Marshal(nu)
  if je != nil{
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("error while encoding json: %v", je)))
  }
  w.Write(jsonResponse)
} 
