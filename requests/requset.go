package requests

type Request struct {
	Name  string `json:"name"`
	Type string `json:"type"`
	Picture	string `json:"picture"`
	Detail string `json:"detail"`
}