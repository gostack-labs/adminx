package api

import (
	"errors"
	"net/http"

	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/utils"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type listApiRequest struct {
	Title    string `json:"title"`
	Groups   int64  `json:"groups" validate:"required"`
	PageID   int32  `json:"page_id" validate:"required,min=1"`
	PageSize int32  `json:"page_size" validate:"required,min=5,max=50"`
}

func (server *Server) listApi(c *bytego.Ctx) error {
	var req listApiRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.ListApiParams{
		Title:      req.Title,
		Groups:     req.Groups,
		Pagelimit:  req.PageSize,
		Pageoffset: (req.PageID - 1) * req.PageSize,
	}
	apiList, err := server.store.ListApi(c.Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, apiList)
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
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	group, err := server.store.GetGroupByID(c.Context(), req.Groups)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.JSON(http.StatusNotFound, errorResponse(err))
		}
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if utils.IsNil(group) {
		return c.JSON(http.StatusNotFound, errorResponse(errors.New("group not found")))
	}
	countArg := db.CountApiByMUTParams{
		Title:  req.Title,
		Url:    req.Url,
		Method: req.Method,
	}
	count, err := server.store.CountApiByMUT(c.Context(), countArg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if count > 0 {
		return c.JSON(http.StatusFound, errorResponse(errors.New("titile url method already exist")))
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
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	group, err := server.store.GetGroupByID(c.Context(), req.Groups)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.JSON(http.StatusNotFound, errorResponse(err))
		}
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if utils.IsNil(group) {
		return c.JSON(http.StatusNotFound, errorResponse(errors.New("group not found")))
	}
	countArg := db.CountApiByMUTParams{
		Title:  req.Title,
		Url:    req.Url,
		Method: req.Method,
	}
	count, err := server.store.CountApiByMUT(c.Context(), countArg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if count > 0 {
		return c.JSON(http.StatusFound, errorResponse(errors.New("titile url method already exist")))
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

type deleteApiRequest struct {
	ID int64 `param:"id" validate:"required"`
}

func (server *Server) deleteApi(c *bytego.Ctx) error {
	var req deleteApiRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	menuApiList, err := server.store.ListMenuApiByApi(c.Context(), []int64{req.ID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if len(menuApiList) > 0 {
		return c.JSON(http.StatusForbidden, errorResponse(errors.New("有菜单绑定了该接口，请解绑后再删除")))
	}

	err = server.store.DeleteApi(c.Context(), []int64{req.ID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "删除成功"})
}

type batchDeleteApiRequest struct {
	IDs []int64 `json:"ids" validate:"required"`
}

func (server *Server) batchDeleteApi(c *bytego.Ctx) error {
	var req batchDeleteApiRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	menuApiList, err := server.store.ListMenuApiByApi(c.Context(), req.IDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if len(menuApiList) > 0 {
		return c.JSON(http.StatusForbidden, errorResponse(errors.New("有菜单绑定了该接口，请解绑后再删除")))
	}

	err = server.store.DeleteApi(c.Context(), req.IDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "删除成功"})
}
