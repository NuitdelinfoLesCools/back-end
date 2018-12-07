package object

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code: 200,
		Data: data,
	}
}

func Fail(err error) *Response {
	return &Response{
		Code:    401,
		Message: err.Error(),
	}
}
