package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)





func  SendResponse[T any](status int, message string, data T, w *http.ResponseWriter){
  response := Response[T]{
    Status: status,
    Message: message,
    Data: data,
  }
  jsonResponse, err := json.Marshal(response)
  if err !=nil{
    fmt.Println("error while encoding json", err)
  }else{
  (*w).Header().Set("Content-Type", "application/json")
  (*w).WriteHeader(status)
  (*w).Write(jsonResponse)
  }
}
