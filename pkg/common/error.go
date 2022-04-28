package common

//CommonError
type CommonError struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (e CommonError) ErrCode() int {
	return e.Code
}
func (e CommonError) Error() string {
	return e.Msg
}

func NewCommonError(code int, msg string, data ...interface{}) *CommonError {
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	return &CommonError{
		Code: code,
		Msg:  msg,
		Data: d,
	}
}
