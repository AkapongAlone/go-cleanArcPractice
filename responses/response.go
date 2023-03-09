package responses

type RequestError struct {
	Status bool              `json:"status"`
	Err    map[string]string `json:"error"`
}

type NoData struct{}

type PaginationBody struct {
	CurrentPage	int		`json:"current_page"`
	Items []ItemBody	`json:"items"`
	NextPage  int		`json:"next_page"`
	PreviousPage int	`json:"previous_page"`
	SizePerPage int		`json:"size_per_page"`
	TotalItems int		`json:"total_item"`
	TotalPages int		`json:"total_page"`

}

type ItemBody struct {
	Created_at string	`json:"created_at"`
	Created_by int `default:"0" json:"created_by"`
	ID 	int			`json:"id"`
	Name	string `json:"name"`
	Picture string	`json:"picture"`
	Detail	string	`json:"detail"`
	Type 	string	`json:"type"`
	Status int 		`default:"0" json:"status"`
	Updated_at string	`json:"updated_at"`
	Updated_by int  `default:"0" json:"updated_by"`
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

type Beer struct {
	ID        int      `json:"id"` 
	Name      string   `json:"name"` 
	Type      string  `json:"type"`  
	Picture   string  `json:"picture"`  
	Detail    string  `json:"detail"`  
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}