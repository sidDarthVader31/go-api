package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var database *gorm.DB
func ConnectDb()(*gorm.DB, error){
  db, err := gorm.Open(postgres.New(postgres.Config{
  DSN: "host=localhost user=sid password=postgresPW dbname=sample-app port= 5455",
  PreferSimpleProtocol: true,
}), &gorm.Config{})
  if err != nil{
    fmt.Println("error while connecting to db")
    return nil, err
  }
  fmt.Println("db value:", db)
  fmt.Println("database value:", database)
  database = db
  return db, nil
}


