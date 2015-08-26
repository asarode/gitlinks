package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  mainRouter := mux.NewRouter()
  v1 := mainRouter.PathPrefix("api/v1").SubRouter()

  v1.HandleFunc("/{username}/projects", readProjects)
    .Methods("GET")

  v1.HandleFunc("{username}/projects", createProjects)
    .Methods("POST")

  v1.HandleFunc("/{username}/projects", updateProjects)
    .Methods("PUT")
}

func readProjects(res http.ResponseWriter, req *http.Request) {
  vars = mux.Vars(req)
  username = vars["username"]
}