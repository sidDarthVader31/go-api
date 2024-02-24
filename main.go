package main

import (
	"fmt"
	"go-api/routes"
	"net/http"

	"github.com/spf13/viper"
)


func main(){

  //initialize env files
  initEnv()
  router := routes.InitRouter();

  //start http server
  http.ListenAndServe(getPort(), router)

  fmt.Println("server started successfully at port", getPort())

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

