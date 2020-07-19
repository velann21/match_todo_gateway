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

func ValidateRegistrationRequest(req *pf.UserRegistrationRequest)error{
	if req.EmailId == ""{
		return InvalidRequest
	}

	if req.Password == ""{
		return InvalidRequest
	}
	return nil
}
