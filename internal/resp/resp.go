package resp

import (
	"net/http"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gostack-labs/adminx/internal/code"
	"github.com/gostack-labs/adminx/pkg/errors"
	"github.com/gostack-labs/bytego"
)

var _ ResultError = (*resultError)(nil)

type ResultError interface {
	// i 为了避免被其他包实现
	i()

	// WithError 设置错误信息
	WithError(err error) ResultError

	// BusinessCode 获取业务码
	BusinessCode() int

	// HTTPCode 获取 HTTP 状态码
	HTTPCode() int

	// Message 获取错误描述
	Error() string

	// StackError 获取带堆栈的错误信息
	StackError() error

	JSON(*bytego.Ctx) error

	AbortJSON(*bytego.Ctx) error
}

type resultError struct {
	httpCode   int         // HTTP 状态码
	Code       int         `json:"code"` // 业务码
	Message    string      `json:"msg"`  // 错误描述
	Detail     interface{} `json:"detail,omitempty"`
	stackError error       // 含有堆栈信息的错误
}

func BadRequestJSON(err error, c *bytego.Ctx) error {
	tv, _ := c.Get("transKey")
	t := tv.(ut.Translator)
	if errs, ok := err.(validator.ValidationErrors); ok {
		rsp := make(bytego.Map)
		for field, terr := range errs.Translate(t) {
			rsp[field[strings.Index(field, ".")+1:]] = terr
		}
		return c.JSON(http.StatusBadRequest, &resultError{
			httpCode: http.StatusBadRequest,
			Code:     code.ParamBindError,
			Message:  code.Text(code.ParamBindError, t.Locale()),
			Detail:   rsp,
		})
	}
	return c.JSON(http.StatusBadRequest, &resultError{
		httpCode: http.StatusBadRequest,
		Code:     code.ParamBindError,
		Message:  code.Text(code.ParamBindError, t.Locale()),
	})
}

func Fail(httpCode, businessCode int) ResultError {
	return &resultError{
		httpCode: httpCode,
		Code:     businessCode,
	}
}

func (e *resultError) i() {}

func (e *resultError) WithError(err error) ResultError {
	e.stackError = errors.WithStack(err)
	return e
}

func (e *resultError) HTTPCode() int {
	return e.httpCode
}

func (e *resultError) BusinessCode() int {
	return e.Code
}

func (e *resultError) Error() string {
	return e.Message
}

func (e *resultError) StackError() error {
	return e.stackError
}

func (e *resultError) JSON(c *bytego.Ctx) error {
	tv, _ := c.Get("transKey")
	t := tv.(ut.Translator)
	e.Message = code.Text(e.Code, t.Locale())
	return c.JSON(e.httpCode, e)
}

func (e *resultError) AbortJSON(c *bytego.Ctx) error {
	tv, _ := c.Get("transKey")
	t := tv.(ut.Translator)
	e.Message = code.Text(e.Code, t.Locale())
	c.Abort()
	return c.JSON(e.httpCode, e)
}

type ResultOK interface {
	JSON(*bytego.Ctx) error
}

type resultOK struct {
	BusinessCode int         `json:"code"`
	Message      string      `json:"msg"`
	Data         interface{} `json:"data,omitempty"`
}

func (o *resultOK) JSON(c *bytego.Ctx) error {
	return c.JSON(http.StatusOK, o)
}

func OK(message string, data ...interface{}) ResultOK {
	var r = new(resultOK)
	r.BusinessCode = 10000
	r.Message = message
	if len(data) > 0 {
		r.Data = data[0]
	}
	return r
}

func DelOK(data ...interface{}) ResultOK {
	var r = new(resultOK)
	r.BusinessCode = 10000
	r.Message = "删除成功"
	if len(data) > 0 {
		r.Data = data[0]
	}
	return r
}

func CreateOK(data ...interface{}) ResultOK {
	var r = new(resultOK)
	r.BusinessCode = 10000
	r.Message = "创建成功"
	if len(data) > 0 {
		r.Data = data[0]
	}
	return r
}

func UpdateOK(data ...interface{}) ResultOK {
	var r = new(resultOK)
	r.BusinessCode = 10000
	r.Message = "修改成功"
	if len(data) > 0 {
		r.Data = data[0]
	}
	return r
}

func OperateOK(data ...interface{}) ResultOK {
	var r = new(resultOK)
	r.BusinessCode = 10000
	r.Message = "操作成功"
	if len(data) > 0 {
		r.Data = data[0]
	}
	return r
}

func GetOK(data ...interface{}) ResultOK {
	var r = new(resultOK)
	r.BusinessCode = 10000
	r.Message = "获取成功"
	if len(data) > 0 {
		r.Data = data[0]
	}
	return r
}
