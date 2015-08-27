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
	"time"
)

func main() {
	mainRouter := mux.NewRouter()
	v1 := mainRouter.PathPrefix("api/v1").SubRouter()

	v1.HandleFunc("/{username}/projects", readProjects).Methods("GET")
	v1.HandleFunc("{username}/projects", createProjects).Methods("POST")
	v1.HandleFunc("/{username}/projects", updateProjects).Methods("PUT")

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

func newUser(username) User {
	return User{
		Created:  time.Now().UnixNano(),
		Username: username,
	}
}

func newProject(owner, repoUri, summary) Project {
	return Project{
		Owner:   owner,
		RepoUri: repoUri,
		Summary: summary,
	}
}

func handleReadProjects(res http.ResponseWriter, req *http.Request) {
	vars = mux.Vars(req)
	username = vars["username"]
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

	err = dbmap.CreateTabledIfNotExists()
	checkError(err, "Create tables failed")
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
