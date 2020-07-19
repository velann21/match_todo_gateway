package controller

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/velann21/match_todo_gateway_srv/pkg/entities/requests"
	"github.com/velann21/match_todo_gateway_srv/pkg/service"
	"net/http"
	"time"
)

type GatewayController struct {
    Service  service.GatewayService
}

func (gctl *GatewayController) UsersController(req *http.Request, res http.ResponseWriter){
	urr, err := requests.PopulateRegistrationApi(req.Body)
	if err != nil{
		logrus.WithError(err).Error("failed to PopulateRegistrationApi()")
		return
	}
	err = requests.ValidateRegistrationRequest(urr)
	if err != nil{
		logrus.WithError(err).Error("failed to PopulateRegistrationApi()")
		return
	}
	parentctx, cancel := context.WithTimeout(req.Context(),  10*time.Second)
	defer cancel()
	apiResp := gctl.Service.RegistrationService(parentctx, urr)


}


