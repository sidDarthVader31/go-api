package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb()(*gorm.DB, error){
  db, err := gorm.Open(postgres.New(postgres.Config{
  DSN: "host=localhost user=sid password=postgresPW dbname=sample-app port= 5455",
  PreferSimpleProtocol: true,
}), &gorm.Config{})
  if err != nil{
    fmt.Println("error while connecting to db")
    return nil, err
  }
  return db, nil
}


