package models

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var database *gorm.DB
func ConnectDb()(*gorm.DB, error){
  host := os.Getenv("HOST")
  port := os.Getenv("DB_PORT")
  user := os.Getenv("USER")
  password := os.Getenv("PASSWORD")
  dbname := os.Getenv("DB_NAME")
  fmt.Printf("host=%s port=%s user=%s password=%s dbname=%s\n", host, port, user, password, dbname)
  db, err := gorm.Open(postgres.New(postgres.Config{ 
  DSN:  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s\n", host, port, user, password, dbname),
  PreferSimpleProtocol: true,
  }), &gorm.Config{})
  if err != nil{
    fmt.Println("error while connecting to db", err)
    return nil, err
  }
  database = db
  return db, nil
}


