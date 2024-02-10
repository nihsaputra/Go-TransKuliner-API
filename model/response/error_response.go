package response

type ErrorResponse struct {
	Code    uint        `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}
