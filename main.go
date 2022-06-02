package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shravan/workday/config"
	userMysql "github.com/shravan/workday/repository/mysql"
	"github.com/shravan/workday/routes"
)

func main() {
	db, err := config.GetDB()
	if err != nil {
		log.Fatal("Unable to connect database", err)
	}
	userRepository := userMysql.NewMysqlUserReposity(db)
	userService := routes.NewUserService(userRepository)
	router := mux.NewRouter()
	routes.NewUserHandler(router, userService)
	http.ListenAndServe(":8080", router)
}
