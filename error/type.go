package error

type DefaultResponse struct {
	DefaultMessage
	Message interface{} `json:"message"`
}

type DefaultErrorResponse struct {
	DefaultMessage
	Error DefaultError `json:"error"`
}

type DefaultError struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

type DefaultMessage struct {
	Type string `json:"type"`
}
