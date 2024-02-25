package users

import (
	"go-api/models"
)



func createUserService(u *models.User, ) error {
  _, err := u.CreateAndGetUser()
  if err!= nil{
    return nil
  }
  return nil
}
