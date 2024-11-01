package users

import (
	"encoding/json"
	"fmt"
	"go-api/common"
	"go-api/models"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, ){
  var nu models.User
  err := json.NewDecoder(r.Body).Decode(&nu)
  if err != nil{
    fmt.Println("error while parsing json:", err)
    common.SendResponse[string](http.StatusBadRequest, " Invalid input",fmt.Sprintf("error:%v", err), &w)
  }
  status, message, _ := CreateUserService(&nu)
  //send ok response 
  common.SendResponse[models.User](status, message, nu, &w)
} 
