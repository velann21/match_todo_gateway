package service

import (
	"context"
	pf "github.com/velann21/todo-commonlib/proto_files/resource_manager"
)

type ResourceManagerService interface {
	CreateCluster(ctx context.Context, request *pf.CreateClusterRequest)(*pf.CreateClusterResponse, error)
}

type ResourceManagerServiceImpl struct {
	ResourceManagerGrpcClient pf.ResourceManagerServiceClient
}

func (rm *ResourceManagerServiceImpl) CreateCluster(ctx context.Context, request *pf.CreateClusterRequest)(*pf.CreateClusterResponse, error){
	resp, err := rm.ResourceManagerGrpcClient.CreateCluster(ctx, request)
	if err != nil{
		return nil, err
	}
	return resp, nil
}
