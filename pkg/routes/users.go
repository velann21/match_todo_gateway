package routes

import (
	"github.com/gorilla/mux"
	"github.com/velann21/match_todo_gateway_srv/pkg/controller"
	"github.com/velann21/match_todo_gateway_srv/pkg/service"
	pf "github.com/velann21/todo-commonlib/proto_files/users_srv"
)

type UsersRoutes struct {
	UsersServiceGrpcClient pf.UserManagementServiceClient
}

func (i *UsersRoutes) UsersRoutes(route *mux.Router){
	gwc := controller.UserServiceController{
		Service:service.New(i.UsersServiceGrpcClient),
	}
	route.PathPrefix("/users/test").HandlerFunc(gwc.TestController).Methods("GET")
	route.PathPrefix("/users/registration").HandlerFunc(gwc.UsersRegistrationController).Methods("POST")
	route.PathPrefix("/users/permission").HandlerFunc(gwc.CreatePermissionsController).Methods("POST")
	route.PathPrefix("/users/roles").HandlerFunc(gwc.CreateRolesController).Methods("POST")
	route.PathPrefix("/users/migration").HandlerFunc(gwc.SqlMigrationController).Methods("POST")
}