package models

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var database *gorm.DB
func ConnectDb()(*gorm.DB, error){
  fmt.Println("database host:", viper.Get("HOST"))
  db, err := gorm.Open(postgres.New(postgres.Config{ 
  DSN:  fmt.Sprintf( "host=%v user=%v password=%v dbname=%v port= %v",viper.Get("HOST"), viper.Get("USER"), viper.Get("PASSWORD"),viper.Get("DB_NAME"), viper.Get("DB_PORT")),
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


