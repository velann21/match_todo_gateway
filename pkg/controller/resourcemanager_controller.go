package controller

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/velann21/match_todo_gateway_srv/pkg/entities/requests"
	"github.com/velann21/match_todo_gateway_srv/pkg/entities/response"
	"github.com/velann21/match_todo_gateway_srv/pkg/service"
	"net/http"
	"time"
)

type ResourceManagerController struct {
	Service  service.ResourceManagerService
	Log *logrus.Logger
}

func (rmc *ResourceManagerController) CreateCluster(res http.ResponseWriter, req *http.Request){
	errResp :=  response.ErrorResponse{}
	successResp := response.SuccessResponse{}
	body, err := requests.PopulateCreateClusterRequest(req.Body)
	if err != nil{
		logrus.WithError(err).Error("failed to PopulateCreateClusterRequest()")
		errResp.HandleError(err, res)
		return
	}
	parentctx, cancel := context.WithTimeout(req.Context(),  10*time.Second)
	defer cancel()
	resp, err := rmc.Service.CreateCluster(parentctx, body)
	if err != nil{
		logrus.WithError(err).Error("failed to CreateCluster()")
		errResp.HandleError(err, res)
		return
	}
	successResp.CreateClusterResponse("12", resp.Success)
	successResp.SuccessResponse(res, http.StatusOK)
}

func (rmc *ResourceManagerController) GenerateEvent(res http.ResponseWriter, req *http.Request) {
	rmc.Log.Info("Inside GenerateEvent")
}