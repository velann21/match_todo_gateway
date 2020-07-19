package service

import (
	"context"
	pf "github.com/velann21/todo-commonlib/proto_files"

)

type GatewayService interface {
	RegistrationService(ctx context.Context, request *pf.UserRegistrationRequest) (*pf.UserRegistrationResponse, error)
}

type GatewayServiceImpl struct {
     UsersServiceGrpcClient pf.UserManagementServiceClient
}

func New(client pf.UserManagementServiceClient) GatewayService {
	return &GatewayServiceImpl{UsersServiceGrpcClient:client}
}

func (srv *GatewayServiceImpl) RegistrationService(ctx context.Context, request *pf.UserRegistrationRequest)(*pf.UserRegistrationResponse, error){
	resp, err := srv.UsersServiceGrpcClient.UserRegistration(ctx, request)
	if err != nil{
		return nil, err
	}
	return resp, nil
}
