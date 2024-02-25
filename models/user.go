package models


type User struct{
  Id uint               `json:"id"`
  Name string           `json:"name"`
  Email string          `json:"email"`
  Password string       `json:"Password"`
  PhoneNumber string    `json:"PhoneNumber"`
  IsActive bool         `json:"isActive"`
}




func (u *User) CreateAndGetUser() (*User, error){
  result :=database.Create(u)
  if result.Error != nil{
    return nil, result.Error
  }
  return u, nil
}

func (u *User) CheckUserExists() (*User, error) {
  result :=database.First(u, "email =?", u.Email)
  if result.Error !=nil{
    return nil, result.Error
  }
  return u, nil
}

func (u *User) DeleteUser() (*User, error){
  result := database.Model(u).Where("id = ?", u.Id).Update("isActive", false)
  if result.Error != nil{
    return nil, result.Error
  }
  return u, nil
}



