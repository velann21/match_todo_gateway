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

type GatewayController struct {
    Service  service.GatewayService
}

func (gctl *GatewayController) UsersController(req *http.Request, res http.ResponseWriter){
	urr, err := requests.PopulateRegistrationApi(req.Body)
	errResp :=  response.ErrorResponse{}
	successResp := response.SuccessResponse{}
	if err != nil{
		logrus.WithError(err).Error("failed to PopulateRegistrationApi()")
		errResp.HandleError(err, res)
	}
	err = requests.ValidateRegistrationRequest(urr)
	if err != nil{
		logrus.WithError(err).Error("failed to ValidateRegistrationRequest()")
		errResp.HandleError(err, res)
	}
	parentctx, cancel := context.WithTimeout(req.Context(),  10*time.Second)
	defer cancel()
	apiResp, err := gctl.Service.RegistrationService(parentctx, urr)
	if err != nil{
		logrus.WithError(err).Error("failed to complete RegistrationService()")
		errResp.HandleError(err, res)
	}
	successResp.UserRegistrationResp(&apiResp.UserId)
	successResp.SuccessResponse(res, http.StatusOK)
}


