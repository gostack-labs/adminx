package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gostack-labs/adminx/internal/middleware/auth"
	"github.com/gostack-labs/adminx/internal/middleware/permission"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/utils"
	"github.com/gostack-labs/adminx/pkg/token"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgx/v4"
)

type listUserRequest struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	PageID   int32  `json:"page_id" validate:"required,min=1"`
	PageSize int32  `json:"page_size" validate:"required,max=50"`
}

func (server *Server) listUser(c *bytego.Ctx) error {
	var req listUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.ListUserParams{
		Username:   req.Username,
		Fullname:   req.FullName,
		Email:      req.Email,
		Phone:      req.Phone,
		Pagelimit:  req.PageSize,
		Pageoffset: (req.PageID - 1) * req.PageSize,
	}
	users, err := server.store.ListUser(c.Context(), arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.JSON(http.StatusNotFound, errorResponse(err))
		}
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, users)
}

func (server *Server) userInfo(c *bytego.Ctx) error {
	var userInfo struct {
		db.User
		Role   []string      `json:"role"`
		Page   []interface{} `json:"page"`
		Button []interface{} `json:"button"`
	}
	payload, exist := c.Get(auth.AuthorizationPayloadKey)
	if exist {
		return c.JSON(http.StatusUnauthorized, errorResponse(errors.New("un authorized")))
	}
	authPayload := payload.(*token.Payload)
	user, err := server.store.GetUser(c.Context(), authPayload.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	userInfo.User = *user
	groups := permission.Enforcer.GetFilteredNamedGroupingPolicy("g", 0, user.Username)
	if len(groups) > 0 {
		roles := make([]string, 0, len(groups))
		for _, g := range groups {
			roles = append(roles, g[1])
		}
		userInfo.Role = roles
	}

	// 查询角色ID
	roleIDs, err := server.store.ListRoleForIDByKeys(c.Context(), userInfo.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	// 查询菜单权限
	argMenu := db.ListRoleMenuForMenuByRolesParams{
		Roles: roleIDs,
		Type:  1,
	}
	menus, err := server.store.ListRoleMenuForMenuByRoles(c.Context(), argMenu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	userInfo.Page, err = server.store.ListMenuForAuthByIDs(c.Context(), menus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	// 查询按钮权限
	argButton := db.ListRoleMenuForMenuByRolesParams{
		Roles: roleIDs,
		Type:  2,
	}
	buttons, err := server.store.ListRoleMenuForMenuByRoles(c.Context(), argButton)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	userInfo.Button, err = server.store.ListMenuForAuthByIDs(c.Context(), buttons)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	return c.JSON(http.StatusOK, userInfo)
}

type userInfoByIDRequest struct {
	Username string `param:"username" validate:"required"`
}

func (server *Server) userInfoByID(c *bytego.Ctx) error {
	var (
		req      userInfoByIDRequest
		userInfo struct {
			db.User
			Role []string
		}
	)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	user, err := server.store.GetUser(c.Context(), req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	userInfo.User = *user
	groups := permission.Enforcer.GetFilteredNamedGroupingPolicy("g", 0, req.Username)
	if len(groups) > 0 {
		roles := make([]string, 0, len(groups))
		for _, g := range groups {
			roles = append(roles, g[1])
		}
		userInfo.Role = roles
	}

	return c.JSON(http.StatusOK, userInfo)
}

type createUserRequest struct {
	Username string   `json:"username" validate:"required,alphanum"`
	Password string   `json:"password" validate:"required,min=6"`
	FullName string   `json:"full_name" validate:"required"`
	Email    string   `json:"email" validate:"required_without=Phone,omitempty,email"`
	Phone    string   `json:"phone" validate:"required_without=Email,omitempty,phone"`
	Role     []string `json:"role"`
}

func (server *Server) createUser(c *bytego.Ctx) error {
	var req createUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if strings.TrimSpace(req.Email) != "" {
		u, err := server.store.GetUserByEmail(c.Context(), req.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		if u != nil {
			return c.JSON(http.StatusForbidden, bytego.Map{"error": "该邮箱已存在！"})
		}
	}

	if strings.TrimSpace(req.Phone) != "" {
		u, err := server.store.GetUserByPhone(c.Context(), req.Phone)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse(err))
		}

		if u != nil {
			return c.JSON(http.StatusForbidden, bytego.Map{"error": "该手机号已存在！"})
		}
	}
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		Phone:          req.Phone,
	}
	var user db.User
	err = server.store.ExecTx(c.Context(), func(q *db.Queries) error {
		user, err := server.store.CreateUser(c.Context(), arg)
		if err != nil {
			return err
		}
		if len(req.Role) > 0 {
			groups := [][]string{}
			for _, role := range req.Role {
				groups = append(groups, []string{user.Username, role})
			}
			_, err = permission.Enforcer.AddNamedGroupingPolicies("g", groups)
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, user)
}

type updateUserRequest struct {
	Username string   `param:"username" validate:"required"`
	FullName string   `json:"full_name" validate:"required"`
	Email    string   `json:"email" validate:"required_without=Phone,omitempty,email"`
	Phone    string   `json:"phone" validate:"required_without=Email,omitempty,phone"`
	Role     []string `json:"role"`
}

func (server *Server) updateUser(c *bytego.Ctx) error {
	var req updateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	err := server.store.ExecTx(c.Context(), func(q *db.Queries) error {
		arg := db.UpdateUserParams{
			Username: req.Username,
			FullName: req.FullName,
			Email:    req.Email,
			Phone:    req.Phone,
		}
		err := server.store.UpdateUser(c.Context(), arg)
		if err != nil {
			return err
		}
		groupMap := make(map[string]struct{})
		for _, role := range req.Role {
			groupMap[role] = struct{}{}
		}
		deleteGroups := make([][]string, 0)
		currentGroups := permission.Enforcer.GetFilteredNamedGroupingPolicy("g", 0, req.Username)
		for _, g := range currentGroups {
			if _, ok := groupMap[g[1]]; !ok {
				deleteGroups = append(deleteGroups, g)
				delete(groupMap, g[1])
			}
		}

		if len(deleteGroups) > 0 {
			_, err = permission.Enforcer.RemoveNamedGroupingPolicies("g", deleteGroups)
			if err != nil {
				return err
			}
		}

		if len(groupMap) > 0 {
			createGroups := make([][]string, 0)
			for k := range groupMap {
				createGroups = append(createGroups, []string{req.Username, k})
			}

			// 保存用户角色关联
			_, err = permission.Enforcer.AddNamedGroupingPolicies("g", createGroups)
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "修改成功"})
}

type deleteUserRequest struct {
	Username string `param:"username" validate:"required"`
}

func (server *Server) deleteUser(c *bytego.Ctx) error {
	var req deleteUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	groups := permission.Enforcer.GetFilteredNamedGroupingPolicy("g", 0, req.Username)
	if len(groups) > 0 {
		return c.JSON(http.StatusFound, errorResponse(errors.New("当前用户关联了其他角色，无法直接删除")))
	}

	err := server.store.DeleteUser(c.Context(), req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "删除成功"})
}
