package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/velann21/match_todo_gateway_srv/pkg/controller"
	"github.com/velann21/match_todo_gateway_srv/pkg/service"
	pf "github.com/velann21/todo-commonlib/proto_files/resource_manager"
)

type ResourceManagerRoutes struct {
	ResourceManagerGrpcClient pf.ResourceManagerServiceClient
	Log *logrus.Logger
}

func (i *ResourceManagerRoutes) ResourceManagerRoutes(route *mux.Router){
	gwc := controller.ResourceManagerController{
		Log: i.Log,
		Service: &service.ResourceManagerServiceImpl{
			ResourceManagerGrpcClient:i.ResourceManagerGrpcClient,
		},
	}
	route.PathPrefix("/resourcemanager/cluster").HandlerFunc(gwc.CreateCluster).Methods("POST")
	route.PathPrefix("/resourcemanager/events").HandlerFunc(gwc.GenerateEvent).Methods("GET")
}