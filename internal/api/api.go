package api

import (
	"net/http"

	"github.com/gostack-labs/adminx/internal/code"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/bytego"
)

type listApiRequest struct {
	Title    string `json:"title"`                                      // 标题
	Groups   int64  `json:"groups" validate:"required"`                 // 接口分组ID
	PageID   int32  `json:"page_id" validate:"required,min=1"`          // 页码
	PageSize int32  `json:"page_size" validate:"required,min=5,max=50"` // 页尺寸
} // 分页获取api列表请求数据

//@title 分页获取api列表接口
//@api get /sys/api
//@group api
//@request listApiRequest
//@response 200 resp.resultOK{businessCode=10000,message="获取成功",data=[]db.Api}
func (server *Server) listApi(c *bytego.Ctx) error {
	var req listApiRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	arg := db.ListApiParams{
		Title:      req.Title,
		Groups:     req.Groups,
		Pagelimit:  req.PageSize,
		Pageoffset: (req.PageID - 1) * req.PageSize,
	}
	apiList, err := server.store.ListApi(c.Context(), arg)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.GetOK(apiList).JSON(c)
}

type createApiRequest struct {
	Title  string `json:"title" validate:"required"`  // 标题
	Url    string `json:"url" validate:"required"`    // 接口地址
	Method string `json:"method" validate:"required"` // 请求方式
	Groups int64  `json:"groups" validate:"required"` // 所属接口分组
	Remark string `json:"remark"`                     // 备注
} // 新增api请求参数

//@title 新增api接口
//@api post /sys/api
//@group api
//@request createApiRequest
//@response 200 resp.resultOK{businessCode=10000,message="创建成功"}
func (server *Server) createApi(c *bytego.Ctx) error {
	var req createApiRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	exist, err := server.store.CheckGroupExist(c.Context(), req.Groups)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if !exist {
		return resp.Fail(http.StatusNotFound, code.ApiGroupNoRowError).JSON(c)
	}
	countArg := db.CountApiByMUTParams{
		Title:  req.Title,
		Url:    req.Url,
		Method: req.Method,
	}
	count, err := server.store.CountApiByMUT(c.Context(), countArg)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if count > 0 {
		return resp.Fail(http.StatusFound, code.ApiExistError).JSON(c)
	}
	arg := db.CreateApiParams{
		Title:  req.Title,
		Url:    req.Url,
		Method: req.Method,
		Groups: req.Groups,
		Remark: &req.Remark,
	}
	err = server.store.CreateApi(c.Context(), arg)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.CreateOK().JSON(c)
}

type updateApiRequest struct {
	ID     int64  `param:"id" validate:"required"`    // 主键ID
	Title  string `json:"title" validate:"required"`  // 标题
	Url    string `json:"url" validate:"required"`    // 接口地址
	Method string `json:"method" validate:"required"` // 请求方式
	Groups int64  `json:"groups" validate:"required"` // 所属接口分组
	Remark string `json:"remark"`                     // 备注
} // 更新api请求参数

//@title 更新api接口
//@api put /sys/api/:id
//@group api
//@request updateApiRequest
//@response 200 resp.resultOK{businessCode=10000,message="修改成功"}
func (server *Server) updateApi(c *bytego.Ctx) error {
	var req updateApiRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	exist, err := server.store.CheckGroupExist(c.Context(), req.Groups)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if !exist {
		return resp.Fail(http.StatusNotFound, code.ApiGroupNoRowError).JSON(c)
	}
	countArg := db.CountApiByMUTParams{
		Title:  req.Title,
		Url:    req.Url,
		Method: req.Method,
	}
	count, err := server.store.CountApiByMUT(c.Context(), countArg)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if count > 0 {
		return resp.Fail(http.StatusFound, code.ApiExistError).JSON(c)
	}
	arg := db.UpdateApiParams{
		ID:     req.ID,
		Title:  req.Title,
		Url:    req.Url,
		Method: req.Method,
		Groups: req.Groups,
		Remark: &req.Remark,
	}
	err = server.store.UpdateApi(c.Context(), arg)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.UpdateOK().JSON(c)
}

type deleteApiRequest struct {
	ID int64 `param:"id" validate:"required"` // 主键ID
} // 删除api请求参数

//@title 删除api接口
//@api delete /sys/api/single/:id
//@group api
//@request deleteApiRequest
//@response 200 resp.resultOK{businessCode=10000,message="删除成功"}
func (server *Server) deleteApi(c *bytego.Ctx) error {
	var req deleteApiRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	menuApiList, err := server.store.ListMenuApiByApi(c.Context(), []int64{req.ID})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if len(menuApiList) > 0 {
		return resp.Fail(http.StatusFound, code.ApiBindMenuExistError).JSON(c)
	}

	err = server.store.DeleteApi(c.Context(), []int64{req.ID})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ApiDeleteError).WithError(err).JSON(c)
	}
	return resp.DelOK().JSON(c)
}

type batchDeleteApiRequest struct {
	IDs []int64 `json:"ids" validate:"required"` // 主键集合
} // 批量删除api请求参数

//@title 批量删除api接口
//@api delete /sys/api/batch
//@group api
//@request batchDeleteApiRequest
//@response 200 resp.resultOK{businessCode=10000,message="删除成功"}
func (server *Server) batchDeleteApi(c *bytego.Ctx) error {
	var req batchDeleteApiRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	menuApiList, err := server.store.ListMenuApiByApi(c.Context(), req.IDs)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if len(menuApiList) > 0 {
		return resp.Fail(http.StatusFound, code.ApiBindMenuExistError).JSON(c)
	}

	err = server.store.DeleteApi(c.Context(), req.IDs)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.DelOK().JSON(c)
}
