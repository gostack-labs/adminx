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
	Key      string `json:"key"`
	PageID   int32  `json:"page_id" validate:"required,min=1"`
	PageSize int32  `json:"page_size" validate:"required,min=5,max=50"`
}

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
	Name   string  `json:"name" validate:"required"`
	Remark *string `json:"remark" validate:"required,omitempty"`
}

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
	ID     int64   `param:"id" validate:"required"`
	Name   string  `json:"name" validate:"required"`
	Remark *string `json:"remark" validate:"required,omitempty"`
}

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
}

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
	IDs []int64 `json:"ids" validate:"required"`
}

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
