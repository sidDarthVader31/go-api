package models

import "gorm.io/gorm"


type User struct{
  Id uint
  Name string
  Email string
  Password string
  PhoneNumber string
  IsActive bool
}




func (u *User) CreateAndGetUser(db *gorm.DB) (*User, error){
  result :=db.Create(u)
  if result.Error != nil{
    return nil, result.Error
  }
  return u, nil
}

func (u *User) CheckUserExists(db *gorm.DB) (*User, error) {
  result :=db.First(u, "email =?", u.Email)
  if result.Error !=nil{
    return nil, result.Error
  }
  return u, nil
}

func (u *User) DeleteUser(db *gorm.DB) (*User, error){
  result := db.Model(u).Where("id = ?", u.Id).Update("isActive", false)
  if result.Error != nil{
    return nil, result.Error
  }
  return u, nil
}



