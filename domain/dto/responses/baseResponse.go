package responses

type BaseResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func (br BaseResponse) Response(status bool, msg string) BaseResponse {
	return BaseResponse{
		Message: msg,
		Status: status,
	}
}