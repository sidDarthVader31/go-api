package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb()(*gorm.DB, error){
  db, err := gorm.Open(postgres.New(postgres.Config{
  DSN: "host=localhost user=gorm password=gorm dbname=gorm port=9920 ",
  PreferSimpleProtocol: true,
}), &gorm.Config{})
  if err != nil{
    fmt.Println("error while connecting to db")
    return nil, err
  }
  return db, nil
}


