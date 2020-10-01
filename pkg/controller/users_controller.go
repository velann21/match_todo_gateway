package controller

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/velann21/match_todo_gateway_srv/pkg/entities/requests"
	"github.com/velann21/match_todo_gateway_srv/pkg/entities/response"
	"github.com/velann21/match_todo_gateway_srv/pkg/service"
	"net/http"
	"time"
)

type UserServiceController struct {
    Service  service.GatewayService
}

func (gctl *UserServiceController) TestController(res http.ResponseWriter, req *http.Request){

	fmt.Println("Starting TestController")
	successResp := response.SuccessResponse{}
	successResp.SuccessResponse(res, http.StatusOK)
	fmt.Println("Done TestController")
	return

}

func (gctl *UserServiceController) UsersRegistrationController(res http.ResponseWriter, req *http.Request){
	urr, err := requests.PopulateRegistrationApi(req.Body)
	errResp :=  response.ErrorResponse{}
	successResp := response.SuccessResponse{}
	if err != nil{
		logrus.WithError(err).Error("failed to PopulateRegistrationApi()")
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

func (gctl *UserServiceController) CreateRolesController(res http.ResponseWriter, req *http.Request){
	rolesReq, err := requests.PopulateCreateRolesApi(req.Body)
	errResp :=  response.ErrorResponse{}
	successResp := response.SuccessResponse{}
	if err != nil{
		logrus.WithError(err).Error("failed to PopulateRegistrationApi()")
		errResp.HandleError(err, res)
	}
	parentctx, cancel := context.WithTimeout(req.Context(),  10*time.Second)
	defer cancel()
	apiResp, err := gctl.Service.CreateRolesService(parentctx, rolesReq)
	if err != nil{
		logrus.WithError(err).Error("failed to complete RegistrationService()")
		errResp.HandleError(err, res)
	}
	successResp.CreateRolesResp(&apiResp.RoleId)
	successResp.SuccessResponse(res, http.StatusOK)
}

func (gctl *UserServiceController) CreatePermissionsController(res http.ResponseWriter, req *http.Request){
	permReq, err := requests.PopulateCreatePermissionApi(req.Body)
	errResp :=  response.ErrorResponse{}
	successResp := response.SuccessResponse{}
	if err != nil{
		logrus.WithError(err).Error("failed to PopulateRegistrationApi()")
		errResp.HandleError(err, res)
	}
	parentctx, cancel := context.WithTimeout(req.Context(),  10*time.Second)
	defer cancel()
	apiResp, err := gctl.Service.CreatePermissionService(parentctx, permReq)
	if err != nil{
		logrus.WithError(err).Error("failed to complete RegistrationService()")
		errResp.HandleError(err, res)
	}
	successResp.CreatePermissionResp(&apiResp.PermissionId)
	successResp.SuccessResponse(res, http.StatusOK)
}

func (gctl *UserServiceController) SqlMigrationController(res http.ResponseWriter, req  *http.Request){
	reqBody, err := requests.PopulateSqlMigrationApi(req.Body)
	errResp :=  response.ErrorResponse{}
	successResp := response.SuccessResponse{Success:true}
	if err != nil{
		logrus.WithError(err).Error("failed to PopulateRegistrationApi()")
		errResp.HandleError(err, res)
		return
	}
	parentctx, cancel := context.WithTimeout(req.Context(),  10*time.Second)
	defer cancel()
	_, err = gctl.Service.SqlMigrationService(parentctx, reqBody)
	if err != nil{
		logrus.WithError(err).Error("failed to complete RegistrationService()")
		errResp.HandleError(err, res)
		return
	}
	successResp.SuccessResponse(res, http.StatusOK)
}




