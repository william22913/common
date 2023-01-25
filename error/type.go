package error

type Formator interface {
	ReformatErrorMessage(err error) DefaultErrorResponse
}

type DefaultResponse struct {
	DefaultMessage
	Message interface{} `json:"message"`
}

type DefaultErrorResponse struct {
	DefaultMessage
	Error DefaultError `json:"error"`
}

type DefaultError struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Reason string `json:"reason"`
}

type DefaultMessage struct {
	Type string `json:"type"`
}

type ErrorParam struct {
	Param       interface{}
	IsConverted bool
}

type Converter func(...interface{}) map[string]ErrorParam

type UnbundledErrorMessages struct {
	status   int
	code     error
	language string
	param    []interface{}
	function Converter
}

func (e UnbundledErrorMessages) Error() string {
	return e.code.Error()
}

func (e *UnbundledErrorMessages) Language(language string) *UnbundledErrorMessages {
	e.language = language
	return e
}

func (e *UnbundledErrorMessages) Param(param ...interface{}) *UnbundledErrorMessages {
	e.param = param
	return e
}
