package routes

import "github.com/gorilla/mux"


func InitRouter() *mux.Router{
  r := mux.NewRouter()
  r.Use(mux.CORSMethodMiddleware(r))

  return r;
}
