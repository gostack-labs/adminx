package api

import (
	"errors"
	"net/http"

	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
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
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.ListApiGroupParams{
		Key:        req.Key,
		Pagelimit:  req.PageSize,
		Pageoffset: (req.PageID - 1) * req.PageSize,
	}
	list, err := server.store.ListApiGroup(c.Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	return c.JSON(http.StatusOK, list)
}

type createApiGroupRequest struct {
	Name   string  `json:"name" validate:"required"`
	Remark *string `json:"remark" validate:"required,omitempty"`
}

func (server *Server) createApiGroup(c *bytego.Ctx) error {
	var req createApiGroupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
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
				return c.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "添加成功"})
}

type updateApiGroupRequest struct {
	ID     int64   `param:"id" validate:"required"`
	Name   string  `json:"name" validate:"required"`
	Remark *string `json:"remark" validate:"required,omitempty"`
}

func (server *Server) updateApiGroup(c *bytego.Ctx) error {
	var req updateApiGroupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
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
				return c.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "修改成功"})
}

type deleteApiGroupRequest struct {
	ID int64 `param:"id" validate:"required"`
}

func (server *Server) deleteApiGroup(c *bytego.Ctx) error {
	var req deleteApiGroupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	apiList, err := server.store.ListApiByGroup(c.Context(), []int64{req.ID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if len(apiList) > 0 {
		return c.JSON(http.StatusForbidden, errorResponse(errors.New("The api group has apis and cannot be deleted directly")))
	}
	err = server.store.DeleteApiGroup(c.Context(), []int64{req.ID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "删除成功"})
}

type batchDeleteApiGroupRequest struct {
	IDs []int64 `json:"ids" validate:"required"`
}

func (server *Server) batchDeleteApiGroup(c *bytego.Ctx) error {
	var req batchDeleteApiGroupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	apiList, err := server.store.ListApiByGroup(c.Context(), req.IDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if len(apiList) > 0 {
		return c.JSON(http.StatusForbidden, errorResponse(errors.New("The api group has apis and cannot be deleted directly")))
	}
	err = server.store.DeleteApiGroup(c.Context(), req.IDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "删除成功"})
}
