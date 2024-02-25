package users

import (
	"encoding/json"
	"fmt"
	"go-api/common"
	"go-api/models"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, ){
  fmt.Println("w:", w)
  fmt.Println("*w:", &w)
  var nu models.User
  err := json.NewDecoder(r.Body).Decode(&nu)
  nu.IsActive = true
  if err != nil{
    fmt.Println("invalid data")
    common.SendResponse[string](http.StatusBadRequest, " Invalid input",fmt.Sprintf("error:%v", err), &w)
  }
  err1 := createUserService(&nu)
  if err1 != nil{
    common.SendResponse[string](http.StatusInternalServerError, "Something went wrong",fmt.Sprintf("error:%v", err), &w)
  }
  //send ok response 

  fmt.Println("new user:", nu)
  common.SendResponse[models.User](http.StatusOK, "Created", nu, &w)
} 
