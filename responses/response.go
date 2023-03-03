package responses

type RequestError struct {
	Status bool              `json:"status"`
	Err    map[string]string `json:"error"`
}

type HandleRequestError struct {
	Status bool   `json:"status"`
	Err    string `json:"error"`
}

func ValidateHandleRequest(msg map[string]string) RequestError {
	var request RequestError
	request.Err = msg
	request.Status = false
	return request
}

func HandleRequest(err error) HandleRequestError {
	var request HandleRequestError
	request.Err = err.Error()
	request.Status = false
	return request
}

func SuccessRequest(err error) HandleRequestError {
	var request HandleRequestError
	request.Err = ""
	request.Status = true
	return request
}
