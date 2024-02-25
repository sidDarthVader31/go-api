package main

import (
	"fmt"
	"go-api/models"
	// "go-api/routes"
	// "net/http"

	"github.com/spf13/viper"
)


func main(){

  //initialize env files
  initEnv()
  // router := routes.InitRouter();

  //start http server
  // http.ListenAndServe(getPort(), router)

  fmt.Println("server started successfully at port", getPort())

  //initialize db 

  db, _ := models.ConnectDb()
  //create user 

  newUser := models.User{
    Name :"sid1",
    Email : "siddharthbisht3108@gmail.com",
    Password: "qwerty",
    PhoneNumber: "8218468189",
    IsActive: true,
  }
   data,_ := newUser.CreateAndGetUser(db)
  fmt.Println("return data", data)
}

func initEnv(){
  viper.SetConfigFile(".env")
  err := viper.ReadInConfig()
  if err!= nil {
    fmt.Println("error while reading env file")
  }

  fmt.Println("successfully parsed env file")
}


func getPort() string{
  return  fmt.Sprintf(":%s", viper.Get("PORT").(string))
}

