package users

import (
	"encoding/json"
	"fmt"
	"go-api/common"
	"go-api/models"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, ){
  var nu models.User
  err := json.NewDecoder(r.Body).Decode(&nu)
  if err != nil{
    fmt.Println("invalid data")
    w.WriteHeader(http.StatusInternalServerError)
    errorResponse := common.Response[string]{Status: http.StatusBadRequest, Message: "Invalid input", Data: fmt.Sprintf("error: %v", err)}
    errorJson, _ := json.Marshal(errorResponse)
    w.Write(errorJson)
  }
  err1 := createUserService(&nu)
  if err1 != nil{
    w.WriteHeader(http.StatusInternalServerError)

    errorResponse := common.Response[string]{
      Status: http.StatusInternalServerError, 
      Message: "Something Went wrong",
      Data: fmt.Sprintf("error: %v", err),
    }
    errorJson, _ := json.Marshal(errorResponse)
    w.Write(errorJson)
  }
  w.WriteHeader(http.StatusOK)
  response := common.Response[models.User]{
    Status: http.StatusOK,
    Data: nu,
    Message: "Created",
  }
  jsonResponse, je := json.Marshal(response)
  if je != nil{
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("error while encoding json: %v", je)))
  }
  w.Write(jsonResponse)
} 
