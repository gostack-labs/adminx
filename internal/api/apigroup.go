package api

import (
	"errors"
	"net/http"

	"github.com/gostack-labs/adminx/internal/code"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgconn"
)

type listApiGroupRequest struct {
	Key      string `json:"key"`                                        // 查询关键字(接口分组名称、备注模糊查询)
	PageID   int32  `json:"page_id" validate:"required,min=1"`          // 页码
	PageSize int32  `json:"page_size" validate:"required,min=5,max=50"` // 页尺寸
} // 分页获取接口分组列表请求参数

//@title 分页获取接口分组接口
//@api get /sys/api-group
//@group api-group
//@request listApiGroupRequest
//@response 200 resp.resultOK{businesscode=10000,message="获取成功",data=[]db.ApiGroup}
func (server *Server) listApiGroup(c *bytego.Ctx) error {
	var req listApiGroupRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	arg := db.ListApiGroupParams{
		Key:        req.Key,
		Pagelimit:  req.PageSize,
		Pageoffset: (req.PageID - 1) * req.PageSize,
	}
	list, err := server.store.ListApiGroup(c.Context(), arg)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.GetOK(list).JSON(c)
}

type createApiGroupRequest struct {
	Name   string  `json:"name" validate:"required"`             // 接口分组名称
	Remark *string `json:"remark" validate:"required,omitempty"` // 备注
} // 新增接口分组请求参数

//@title 新增接口分组接口
//@api post /sys/api-group
//@group api-group
//@request createApiGroupRequest
//@response 200 resp.resultOK{businesscode=10000,message="创建成功"}
func (server *Server) createApiGroup(c *bytego.Ctx) error {
	var req createApiGroupRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	arg := db.CreateApiGroupParams{
		Name:   req.Name,
		Remark: req.Remark,
	}
	err := server.store.CreateApiGroup(c.Context(), arg)
	if err != nil {
		var pgxerr *pgconn.PgError
		if errors.As(err, &pgxerr) {
			if pgxerr.Code == "23505" {
				return resp.Fail(http.StatusFound, code.ApiGroupExistError).JSON(c)
			}
		}
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.CreateOK().JSON(c)
}

type updateApiGroupRequest struct {
	ID     int64   `param:"id" validate:"required"`              // 主键ID
	Name   string  `json:"name" validate:"required"`             // 接口分组名称
	Remark *string `json:"remark" validate:"required,omitempty"` // 备注
} // 更新接口分组请求参数

//@title 更新接口分组接口
//@api put /sys/api-group/:id
//@group api-group
//@request updateApiGroupRequest
//@response 200 resp.resultOK{businesscode=10000,message="修改成功"}
func (server *Server) updateApiGroup(c *bytego.Ctx) error {
	var req updateApiGroupRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	arg := db.UpdateApiGroupParams{
		Name:   req.Name,
		Remark: req.Remark,
		ID:     req.ID,
	}
	err := server.store.UpdateApiGroup(c.Context(), arg)
	if err != nil {
		var pgxerr *pgconn.PgError
		if errors.As(err, &pgxerr) {
			if pgxerr.Code == "23505" {
				return resp.Fail(http.StatusFound, code.ApiGroupExistError).JSON(c)
			}
		}
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.UpdateOK().JSON(c)
}

type deleteApiGroupRequest struct {
	ID int64 `param:"id" validate:"required"`
} // 删除接口分组请求参数

//@title 删除接口分组接口
//@api delete /sys/api-group/single/:id
//@group api-group
//@request deleteApiGroupRequest
//@response 200 resp.resultOK{businesscode=10000,message="删除成功"}
func (server *Server) deleteApiGroup(c *bytego.Ctx) error {
	var req deleteApiGroupRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	apiList, err := server.store.ListApiByGroup(c.Context(), []int64{req.ID})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if len(apiList) > 0 {
		return resp.Fail(http.StatusFound, code.ApiGroupHasApiError).JSON(c)
	}
	err = server.store.DeleteApiGroup(c.Context(), []int64{req.ID})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.DelOK().JSON(c)
}

type batchDeleteApiGroupRequest struct {
	IDs []int64 `json:"ids" validate:"required"` // 主键ID集合
} // 批量删除接口分组请求参数

//@title 新增接口分组接口
//@api delete /sys/api-group/batch
//@group api-group
//@request batchDeleteApiGroupRequest
//@response 200 resp.resultOK{businesscode=10000,message="删除成功"}
func (server *Server) batchDeleteApiGroup(c *bytego.Ctx) error {
	var req batchDeleteApiGroupRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	apiList, err := server.store.ListApiByGroup(c.Context(), req.IDs)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if len(apiList) > 0 {
		return resp.Fail(http.StatusFound, code.ApiGroupHasApiError).JSON(c)
	}
	err = server.store.DeleteApiGroup(c.Context(), req.IDs)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.DelOK().JSON(c)
}
