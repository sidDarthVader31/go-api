package models


type User struct{
  Id uint
  Name string
  Email string
  Password string
  PhoneNumber string
  IsActive bool
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



