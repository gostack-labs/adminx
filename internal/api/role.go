package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gostack-labs/adminx/internal/middleware/permission"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgconn"
	"github.com/spf13/cast"
)

type listRoleRequest struct {
	PageID   int32  `json:"page_id" validate:"required,min=1"`
	PageSize int32  `json:"page_size" validate:"required,max=50"`
	Name     string `json:"name"`
	Key      string `json:"key"`
}

func (server *Server) listRole(c *bytego.Ctx) error {
	var req listRoleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListRoleParams{
		Pagelimit:  req.PageSize,
		Pageoffset: (req.PageID - 1) * req.PageSize,
		Name:       req.Name,
		Key:        req.Key,
	}

	listRole, err := server.store.ListRole(c.Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	return c.JSON(http.StatusOK, listRole)
}

type createRoleRequest struct {
	Name      string  `json:"name" validate:"required"`
	IsDisable bool    `json:"is_disable"`
	Key       string  `json:"key" validate:"required"`
	Sort      int32   `json:"sort"`
	Remark    *string `json:"remark"`
}

func (server *Server) createRole(c *bytego.Ctx) error {
	var req createRoleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateRoleParams{
		Name:      req.Name,
		IsDisable: req.IsDisable,
		Key:       req.Key,
		Sort:      req.Sort,
		Remark:    req.Remark,
	}
	err := server.store.CreateRole(c.Context(), arg)
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

type updateRoleRequest struct {
	ID        int64   `param:"id" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	IsDisable bool    `json:"is_disable"`
	Key       string  `json:"key" validate:"required"`
	Sort      int32   `json:"sort"`
	Remark    *string `json:"remark"`
}

func (server *Server) updateRole(c *bytego.Ctx) error {
	var req updateRoleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateRoleParams{
		ID:        req.ID,
		Name:      req.Name,
		IsDisable: req.IsDisable,
		Key:       req.Key,
		Sort:      req.Sort,
		Remark:    req.Remark,
	}
	err := server.store.UpdateRole(c.Context(), arg)
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

type deleteRoleRequest struct {
	ID int64 `param:"id" validate:"required"`
}

func (server *Server) deleteRole(c *bytego.Ctx) error {
	var req deleteRoleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// 获取角色绑定的用户
	roles, err := server.store.GetRoleKeyByIDs(c.Context(), []int64{req.ID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	groups := permission.Enforcer.GetFilteredNamedGroupingPolicy("g", 1, roles...)
	if len(groups) > 0 {
		return c.JSON(http.StatusFound, errorResponse(errors.New("角色被其他用户关联，无法删除")))
	}

	// 获取角色绑定的菜单
	countRoleMenu, err := server.store.CountRoleMenuByRole(c.Context(), []int64{req.ID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if countRoleMenu > 0 {
		return c.JSON(http.StatusFound, errorResponse(errors.New("角色存在与菜单的关联，无法删除")))
	}
	err = server.store.DeleteRole(c.Context(), []int64{req.ID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "删除成功"})
}

type batchDeleteRoleRequest struct {
	IDs []int64 `json:"ids" validate:"required"`
}

func (server *Server) batchDeleteRole(c *bytego.Ctx) error {
	var req batchDeleteRoleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// 获取角色绑定的用户
	roles, err := server.store.GetRoleKeyByIDs(c.Context(), req.IDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	groups := permission.Enforcer.GetFilteredNamedGroupingPolicy("g", 1, roles...)
	if len(groups) > 0 {
		return c.JSON(http.StatusFound, errorResponse(errors.New("角色被其他用户关联，无法删除")))
	}

	// 获取角色绑定的菜单
	countRoleMenu, err := server.store.CountRoleMenuByRole(c.Context(), req.IDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if countRoleMenu > 0 {
		return c.JSON(http.StatusFound, errorResponse(errors.New("角色存在与菜单的关联，无法删除")))
	}
	err = server.store.DeleteRole(c.Context(), req.IDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "删除成功"})
}

type updateRolePermissionRequest struct {
	ID        int64                     `param:"id" validate:"required"`
	RoleMenus []db.CreateRoleMenuParams `json:"role_menus"`
}

func (server *Server) updateRolePermission(c *bytego.Ctx) error {
	var (
		req            updateRolePermissionRequest
		createRoleMenu []db.CreateRoleMenuParams
		errs           []error
	)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// 查询角色对应的菜单列表
	rm, err := server.store.ListRoleMenuByRole(c.Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if len(rm) > 0 {
		oldRoleMenuMap := make(map[string]int64)
		newRoleMenuMap := make(map[string]struct{})
		for _, oldRoleMenu := range rm {
			oldRoleMenuMap[fmt.Sprintf("%d-%d-%d", oldRoleMenu.Menu, oldRoleMenu.Menu, oldRoleMenu.Type)] = oldRoleMenu.ID
		}
		for _, newRoleMenu := range req.RoleMenus {
			roleMenuKey := fmt.Sprintf("%d-%d-%d", newRoleMenu.Role, newRoleMenu.Menu, newRoleMenu.Type)
			if _, ok := oldRoleMenuMap[roleMenuKey]; ok {
				delete(oldRoleMenuMap, roleMenuKey)
			} else {
				roleMenuSlice := strings.Split(roleMenuKey, "-")
				roleIdTmp := cast.ToInt64(roleMenuSlice[0])
				menuIdTmp := cast.ToInt64(roleMenuSlice[1])
				typeTmp := cast.ToInt32(roleMenuSlice[2])
				createRoleMenu = append(createRoleMenu, db.CreateRoleMenuParams{
					Role: roleIdTmp,
					Menu: menuIdTmp,
					Type: typeTmp,
				})
			}
			newRoleMenuMap[roleMenuKey] = struct{}{}
		}
		if len(createRoleMenu) > 0 {
			crm := server.store.CreateRoleMenu(c.Context(), createRoleMenu)
			defer crm.Close()
			crm.Exec(func(i int, err error) {
				errs = append(errs, err)
			})
			for _, v := range errs {
				if v != nil {
					return c.JSON(http.StatusInternalServerError, errorResponse(err))
				}
			}
		}
		// 需要删除的节点
		deleteRoleMenuList := make([]int64, 0, len(oldRoleMenuMap))
		for _, v := range oldRoleMenuMap {
			deleteRoleMenuList = append(deleteRoleMenuList, v)
		}
		if len(deleteRoleMenuList) > 0 {
			err := server.store.DeleteRoleMenu(c.Context(), deleteRoleMenuList)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, errorResponse(err))
			}
		}
	} else {
		crm := server.store.CreateRoleMenu(c.Context(), req.RoleMenus)
		defer crm.Close()
		crm.Exec(func(i int, err error) {
			errs[i] = err
		})
		for _, v := range errs {
			if v != nil {
				return c.JSON(http.StatusInternalServerError, errorResponse(v))
			}
		}
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "创建成功"})
}

type getRolePermissionRequest struct {
	ID int64 `param:"id" validate:"required"`
}

func (server *Server) getRolePermission(c *bytego.Ctx) error {
	var req getRolePermissionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	// 查询所有目录ID
	parentMenuIDs, err := server.store.ListMenuForParent(c.Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	// 查询当前角色的菜单
	arg := db.ListRoleMenuForMenuParams{
		Role:         req.ID,
		Excludemenus: parentMenuIDs,
	}
	roleMenus, err := server.store.ListRoleMenuForMenu(c.Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	// 查询当前角色的按钮
	buttons, err := server.store.ListRoleMenuForButton(c.Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	menuButtons, err := server.store.ListMenuForParentIDByID(c.Context(), buttons)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	roleButtons := make(map[int64][]int64)
	for _, v := range menuButtons {
		roleButtons[v.Parent] = append(roleButtons[v.Parent], v.ID)
	}
	return c.JSON(http.StatusOK, bytego.Map{
		"menu":   roleMenus,
		"button": roleButtons,
	})
}

type roleApiPermissionRequest struct {
	ID   int64   `param:"id" validate:"required"`
	Type *int    `json:"type" validate:"required,oneof=0 1"` // 0:解除api权限 1:绑定api权限
	Api  []int64 `json:"api" validate:"required"`
}

func (server *Server) roleApiPermission(c *bytego.Ctx) error {
	var req roleApiPermissionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	role, err := server.store.ListRoleByID(c.Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	apis, err := server.store.ListApiByIDs(c.Context(), req.Api)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	var groups [][]string
	for _, api := range apis {
		groups = append(groups, []string{role.Key, api.Url, api.Method})
	}

	if *req.Type == 1 {
		_, err = permission.Enforcer.AddNamedPolicies("p", groups)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse(err))
		}
	} else if *req.Type == 0 {
		_, err := permission.Enforcer.RemoveNamedPolicies("p", groups)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse(err))
		}
	}
	return c.JSON(http.StatusOK, bytego.Map{"success": true, "message": "操作成功"})
}

type getRoleApiRequest struct {
	ID   int64 `param:"id" validate:"required"`
	Menu int64 `json:"menu" validate:"required"`
}

func (server *Server) getRoleApi(c *bytego.Ctx) error {
	var (
		req  getRoleApiRequest
		errs []error
		ids  []int64
		apis []int64
	)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	role, err := server.store.ListRoleByID(c.Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	roleApis := permission.Enforcer.GetFilteredNamedPolicy("p", 0, role.Key)
	if len(roleApis) > 0 {
		menuApiIds, err := server.store.ListMenuApiForApiByMenu(c.Context(), req.Menu)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		args := []db.ListApiBatchParams{}
		for _, p := range roleApis {
			args = append(args, db.ListApiBatchParams{Url: p[1], Method: p[2]})
		}
		labbr := server.store.ListApiBatch(c.Context(), args)
		defer labbr.Close()
		labbr.Query(func(i int, idList []int64, err error) {
			ids = append(ids, idList...)
			errs = append(errs, err)
		})
		for _, ma := range menuApiIds {
			for _, id := range ids {
				if ma == id {
					apis = append(apis, id)
				}
			}
		}
	}
	if len(apis) == 0 {
		apis = []int64{}
	}
	return c.JSON(http.StatusOK, apis)
}
