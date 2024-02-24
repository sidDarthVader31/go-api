package models

import "gorm.io/gorm"


type User struct{
  Id uint
  Name string
  Email string
  password string
  phoneNumber string
}




func (u *User) createAndGetUser(db *gorm.DB) (*User, error){
   db.Create(u)
  return u, nil
}




