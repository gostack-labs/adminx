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
//@response 200 resp.resultOK{code=10000,msg="获取成功",data=[]*db.Api}
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
	Title  string `json:"title" validate:"required"`
	Url    string `json:"url" validate:"required"`
	Method string `json:"method" validate:"required"`
	Groups int64  `json:"groups" validate:"required"`
	Remark string `json:"remark"`
}

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
	ID     int64  `param:"id" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Url    string `json:"url" validate:"required"`
	Method string `json:"method" validate:"required"`
	Groups int64  `json:"groups" validate:"required"`
	Remark string `json:"remark"`
}

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
	ID int64 `param:"id" validate:"required"`
}

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
	IDs []int64 `json:"ids" validate:"required"`
}

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
