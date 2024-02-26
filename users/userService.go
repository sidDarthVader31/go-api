package users

import (
	"fmt"
	"go-api/models"
	"net/http"
)



func CreateUserService(u *models.User ) (int, string, models.User) {
  //check user exis
  userExists := doesUserExist(u)
  if userExists == true{
    return http.StatusBadRequest, "User already exists", *u
  }
  _, err := u.CreateAndGetUser()
  if err!= nil{
    return http.StatusOK, "User created", *u
  }
  return http.StatusInternalServerError, "Someting went wrong", *u
}

func doesUserExist(u *models.User) bool {
  data, err := u.CheckUserExists()
  fmt.Println("data:", data)
  fmt.Println("error:", err)
  if data == nil {
    return false
  }
  return true
} 
