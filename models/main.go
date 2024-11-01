package models

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var database *gorm.DB
func ConnectDb()(*gorm.DB, error){
  fmt.Println("database host:", viper.Get("HOST"))

connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    os.Getenv("HOST"),    // "postgres-service"
    os.Getenv("DB_PORT"),    // "5432"
    os.Getenv("USER"),
    os.Getenv("PASSWORD"),
    os.Getenv("DB_NAME"))
  fmt.Println("conn str:", connStr)
  db, err := gorm.Open(postgres.New(postgres.Config{ 
  DSN:  connStr,
  PreferSimpleProtocol: true,
}), &gorm.Config{})
  if err != nil{
    fmt.Println("error while connecting to db", err)
    return nil, err
  }
  fmt.Println("db value:", db)
  fmt.Println("database value:", database)
  database = db
  return db, nil
}


