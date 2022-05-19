package api

import (
	"go/token"
	"log"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	ut "github.com/go-playground/universal-translator"
	"github.com/gostack-labs/adminx/internal/code"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/bytego"
	"github.com/spf13/cast"
)

const minBusinessCode = 20000

type codes struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误描述
}

type codeResponse struct {
	SystemCodes   []codes // 系统错误
	BusinessCodes []codes // 业务错误
} // 获取错误列表

//@title 获取错误列表接口
//@api get /code
//@group basic
//@response 200 resp.resultOK{businesscode=10000,message="获取成功",data=codeResponse}
func (server *Server) code(c *bytego.Ctx) error {
	tv, _ := c.Get("transKey")
	t := tv.(ut.Translator)
	parsedFile, err := decorator.Parse(code.ByteCodeFile)
	if err != nil {
		log.Fatalf("parsing code.go: %s:%s\n", "ByteCodeFile", err)
	}

	var (
		systemCodes   []codes
		businessCodes []codes
	)

	dst.Inspect(parsedFile, func(n dst.Node) bool {
		decl, ok := n.(*dst.GenDecl)
		if !ok || decl.Tok != token.CONST {
			return true
		}

		for _, spec := range decl.Specs {
			valueSpec, _ok := spec.(*dst.ValueSpec)
			if !_ok {
				continue
			}

			codeInt := cast.ToInt(valueSpec.Values[0].(*dst.BasicLit).Value)

			if codeInt < minBusinessCode {
				systemCodes = append(systemCodes, codes{
					Code:    codeInt,
					Message: code.Text(codeInt, t.Locale()),
				})
			} else {
				businessCodes = append(businessCodes, codes{
					Code:    codeInt,
					Message: code.Text(codeInt, t.Locale()),
				})
			}
		}
		return true
	})
	return resp.GetOK(codeResponse{SystemCodes: systemCodes, BusinessCodes: businessCodes}).JSON(c)
}
