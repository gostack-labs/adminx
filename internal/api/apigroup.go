package api

import (
	"errors"
	"net/http"

	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgconn"
)

type listApiGroupRequest struct {
	Key string `json:"key"`
}

func (server *Server) listApiGroup(c *bytego.Ctx) error {
	var req listApiGroupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	list, err := server.store.ListApiGroup(c.Context(), req.Key)
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
	return c.JSON(http.StatusOK, struct{}{})
}
