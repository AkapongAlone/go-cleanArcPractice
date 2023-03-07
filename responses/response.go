package responses

type RequestError struct {
	Status bool              `json:"status"`
	Err    map[string]string `json:"error"`
}

type NoData struct{}

type PaginationBody struct {
	CurrentPage	int
	Items []ItemBody
	NextPage  int
	PreviousPage int
	SizePerPage int
	TotalItems int
	TotalPages int

}

type ItemBody struct {
	Created_at string
	Created_by int `default:"0"`
	ID 	int
	Name	string
	Picture string
	Detail	string
	Type 	string
	Status int 		`default:"0"`
	Updated_at string
	Updated_by int  `default:"0"`
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

type NoDataResponse struct {
	Status bool   `json:"status" example:"true" extensions:"x-order=0"`
	Code   int    `json:"code" extensions:"x-order=1"`
	Data   NoData `json:"data" extensions:"x-order=1"`
}