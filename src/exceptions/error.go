package exceptions

import "time"

type Response struct {
	ErrorCode int    `json:"error_code"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

func (err Response) Create(errorCode int, message string) Response {
	return Response{
		ErrorCode: errorCode,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339Nano),
	}
}
