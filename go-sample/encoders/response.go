package encoders

type Response struct {
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}
