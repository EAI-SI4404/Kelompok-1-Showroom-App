package services

type ErrorMessage struct {
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
}

func (e *ErrorMessage) Error() string {
	return e.Message
}