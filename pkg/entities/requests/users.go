package requests

import (
	"encoding/json"
	"errors"
	"io"
	pf "github.com/velann21/todo-commonlib/proto_files"
)

var (
	InvalidRequest = errors.New("Invalid request")
)
func PopulateRegistrationApi(body io.Reader)(*pf.UserRegistrationRequest, error){
	req := &pf.UserRegistrationRequest{}
	decode := json.NewDecoder(body)
	err := decode.Decode(req)
	if err != nil{
		return nil, InvalidRequest
	}
	return req, nil
}

func PopulateCreatePermissionApi(body io.Reader)(*pf.CreatePermissionRequest, error){
	req := &pf.CreatePermissionRequest{}
	decode := json.NewDecoder(body)
	err := decode.Decode(req)
	if err != nil{
		return nil, InvalidRequest
	}
	return req, nil
}

func PopulateCreateRolesApi(body io.Reader)(*pf.CreateRoleRequest, error){
	req := &pf.CreateRoleRequest{}
	decode := json.NewDecoder(body)
	err := decode.Decode(req)
	if err != nil{
		return nil, InvalidRequest
	}
	return req, nil
}

func PopulateSqlMigrationApi(body io.Reader)(*pf.SqlMigrationRequest, error){
	req := &pf.SqlMigrationRequest{}
	decode := json.NewDecoder(body)
	err := decode.Decode(req)
	if err != nil{
		return nil, InvalidRequest
	}
	return req, nil
}