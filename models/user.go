package models

import "gorm.io/gorm"


type User struct{
  Id uint
  Name string
  Email string
  password string
  phoneNumber string
  isActive bool
}




func (u *User) createAndGetUser(db *gorm.DB) (*User, error){
  result :=db.Create(u)
  if result.Error != nil{
    return nil, result.Error
  }
  return u, nil
}

func (u *User) checkUserExists(db *gorm.DB) (*User, error) {
  result :=db.First(u, "email =?", u.Email)
  if result.Error !=nil{
    return nil, result.Error
  }
  return u, nil
}

func (u *User) deleteUser(db *gorm.DB) (*User, error){
  result := db.Where("id = ?", u.Id).Update("isActive", false)
  if result.Error != nil{
    return nil, result.Error
  }
  return u, nil
}



