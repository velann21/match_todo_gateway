package service

import (
	"context"
	"github.com/sirupsen/logrus"
	pf "github.com/velann21/todo-commonlib/proto_files/users_srv"

)

type GatewayService interface {
	RegistrationService(ctx context.Context, request *pf.UserRegistrationRequest) (*pf.UserRegistrationResponse, error)
	CreatePermissionService(ctx context.Context, request *pf.CreatePermissionRequest)(*pf.CreatePermissionResponse, error)
	CreateRolesService(ctx context.Context, request *pf.CreateRoleRequest)(*pf.CreateRoleResponse, error)
	SqlMigrationService(ctx context.Context, request *pf.SqlMigrationRequest)(*pf.SqlMigrationResponse, error)
}

type GatewayServiceImpl struct {
     UsersServiceGrpcClient pf.UserManagementServiceClient
}

func New(client pf.UserManagementServiceClient) GatewayService {
	logrus.Debug("Creating user service object")
	return &GatewayServiceImpl{UsersServiceGrpcClient:client}
}

func (srv *GatewayServiceImpl) RegistrationService(ctx context.Context, request *pf.UserRegistrationRequest)(*pf.UserRegistrationResponse, error){
	resp, err := srv.UsersServiceGrpcClient.UserRegistration(ctx, request)
	if err != nil{
		logrus.WithError(err).Error("Grpc call failed UserRegistration")
		return nil, err
	}
	return resp, nil
}

func (srv *GatewayServiceImpl) CreatePermissionService(ctx context.Context, request *pf.CreatePermissionRequest)(*pf.CreatePermissionResponse, error){
	resp, err := srv.UsersServiceGrpcClient.CreatePermission(ctx, request)
	if err != nil{
		logrus.WithError(err).Error("Grpc call failed CreatePermissionService")
		return nil, err
	}
	return resp, nil
}

func (srv *GatewayServiceImpl) SqlMigrationService(ctx context.Context, request *pf.SqlMigrationRequest)(*pf.SqlMigrationResponse, error){
	resp, err := srv.UsersServiceGrpcClient.SqlMigration(ctx, request)
	if err != nil{
		logrus.WithError(err).Error("Grpc call failed SqlMigrationService")
		return nil, err
	}
	return resp, nil
}


func (srv *GatewayServiceImpl) CreateRolesService(ctx context.Context, request *pf.CreateRoleRequest)(*pf.CreateRoleResponse, error){
	resp, err := srv.UsersServiceGrpcClient.CreateRole(ctx, request)
	if err != nil{
		logrus.WithError(err).Error("Grpc call failed CreateRolesService")
		return nil, err
	}
	return resp, nil
}

