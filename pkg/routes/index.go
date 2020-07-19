package routes

import "github.com/gorilla/mux"

type IndexRoutes struct {

}

func ApiRoutes(route mux.Router){
	route.PathPrefix("/api/v1/todo/create").HandlerFunc(NewController(conf).CreateTodoController)
}