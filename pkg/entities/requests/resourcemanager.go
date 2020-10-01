package requests

import (
	"encoding/json"
	pf "github.com/velann21/todo-commonlib/proto_files/resource_manager"
	"io"
)

func PopulateCreateClusterRequest(body io.Reader)(*pf.CreateClusterRequest, error){
	req := &pf.CreateClusterRequest{}
	decode := json.NewDecoder(body)
	err := decode.Decode(req)
	if err != nil{
		return nil, InvalidRequest
	}
	return req, nil
}

