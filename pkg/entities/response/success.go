package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Success bool
	Status string
	Data []map[string]interface{}
}

func (entity *SuccessResponse) UserRegistrationResp(id *string) {
	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	data["id"] = *id
	responseData = append(responseData, data)
	entity.Data = responseData
	metaData := make(map[string]interface{})
	metaData["message"] = "User registered"
}



func (resp *SuccessResponse) SuccessResponse(rw http.ResponseWriter, statusCode int){
	rw.Header().Set("Content-Type", "application/json")

	switch statusCode {
	case http.StatusOK:
		rw.WriteHeader(http.StatusOK)
		resp.Status = http.StatusText(http.StatusOK)
	case http.StatusCreated:
		rw.WriteHeader(http.StatusCreated)
		resp.Status = http.StatusText(http.StatusCreated)
	case http.StatusAccepted:
		rw.WriteHeader(http.StatusAccepted)
		resp.Status = http.StatusText(http.StatusAccepted)
	default:
		rw.WriteHeader(http.StatusOK)
		resp.Status = http.StatusText(http.StatusOK)
	}
	// send response
	_ = json.NewEncoder(rw).Encode(resp)
	return
}


