package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mainRouter := mux.NewRouter()
	v1 := mainRouter.PathPrefix("api/v1").Subrouter()

	v1.HandleFunc("/{username}/projects", handleReadProjects).Methods("GET")
	v1.HandleFunc("{username}/projects", handleCreateProjects).Methods("POST")
	v1.HandleFunc("/{username}/projects", handleUpdateProjects).Methods("PUT")

	// Initialize the DB map
	dbmap := initDb()
	defer dbmap.Db.Close()
}

type Config struct {
	DbUri string
}

type User struct {
	Id       int64  `db:"user_id, primarykey, autoincrement"`
	Created  int64  `db:"created_at"`
	Username string `db:"username"`
}

type Project struct {
	Id      int64  `db:"project_id"`
	Owner   int64  `db:"owner_id"`
	RepoUri string `db:"repo_uri"`
	Summary string `db:"summary"`
}

func newUser(username string) User {
	return User{
		Created:  time.Now().UnixNano(),
		Username: username,
	}
}

func newProject(owner int64, repoUri string, summary string) Project {
	return Project{
		Owner:   owner,
		RepoUri: repoUri,
		Summary: summary,
	}
}

func handleReadProjects(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	fmt.Println(username)
}

func handleCreateProjects(res http.ResponseWriter, req *http.Request) {

}

func handleUpdateProjects(res http.ResponseWriter, req *http.Request) {

}

func initDb() *gorp.DbMap {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	checkError(err, "Failed to decode config file")

	db, err := sql.Open("postgres", config.DbUri)
	checkError(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(User{}, "users")
	dbmap.AddTableWithName(Project{}, "projects")

	err = dbmap.CreateTablesIfNotExists()
	checkError(err, "Create tables failed")

	return dbmap
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
