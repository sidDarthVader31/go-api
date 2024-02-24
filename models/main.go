package models

import "github.com/go-pg/pg/v10"



func ConnectDatabase() *pg.DB{
  db := pg.Connect(&pg.Options{
    Addr:     ":5432",
    User:     "user",
    Password: "pass",
    Database: "db_name",
  })
  return db
}
