package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gostack-labs/adminx/internal/code"
	"github.com/gostack-labs/adminx/internal/middleware/permission"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgconn"
	"github.com/spf13/cast"
)

type listRoleRequest struct {
	PageID   int32  `json:"page_id" validate:"required,min=1"`    // 页码
	PageSize int32  `json:"page_size" validate:"required,max=50"` // 页尺寸
	Name     string `json:"name"`                                 // 名称
	Key      string `json:"key"`                                  // 标识
} // 分页获取角色列表请求参数

//@title 分页获取角色列表接口
//@api get /sys/role
//@group role
//@request listRoleRequest
//@response 200 resp.resultOK{businesscode=10000,message="获取成功",data=[]db.Role}
func (server *Server) listRole(c *bytego.Ctx) error {
	var req listRoleRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	arg := db.ListRoleParams{
		Pagelimit:  req.PageSize,
		Pageoffset: (req.PageID - 1) * req.PageSize,
		Name:       req.Name,
		Key:        req.Key,
	}

	listRole, err := server.store.ListRole(c.Context(), arg)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.GetOK(listRole).JSON(c)
}

type createRoleRequest struct {
	Name      string  `json:"name" validate:"required"` // 名称
	IsDisable bool    `json:"is_disable"`               // 是否禁用
	Key       string  `json:"key" validate:"required"`  // 标识
	Sort      int32   `json:"sort"`                     // 排序
	Remark    *string `json:"remark"`                   // 备注
} // 创建角色请求参数

//@title 创建角色接口
//@api post /sys/role
//@group role
//@request createRoleRequest
//@response 200 resp.resultOK{businesscode=10000,message="创建成功"}
func (server *Server) createRole(c *bytego.Ctx) error {
	var req createRoleRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
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
				return resp.Fail(http.StatusFound, code.RoleExistError).JSON(c)
			}
		}
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.CreateOK().JSON(c)
}

type updateRoleRequest struct {
	ID        int64   `param:"id" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	IsDisable bool    `json:"is_disable"`
	Key       string  `json:"key" validate:"required"`
	Sort      int32   `json:"sort"`
	Remark    *string `json:"remark"`
} // 更新角色请求参数

//@title 更新角色接口
//@api put /sys/role/:id
//@group role
//@request updateRoleRequest
//@response 200 resp.resultOK{businesscode=10000,message="更新成功"}
func (server *Server) updateRole(c *bytego.Ctx) error {
	var req updateRoleRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
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
				return resp.Fail(http.StatusFound, code.ServerError).JSON(c)
			}
		}
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.UpdateOK().JSON(c)
}

type deleteRoleRequest struct {
	ID int64 `param:"id" validate:"required"` // 主键ID
} // 删除角色请求参数

//@title 删除角色接口
//@api delete /sys/role/single/:id
//@group role
//@request deleteRoleRequest
//@response 200 resp.resultOK{businesscode=10000,message="删除成功"}
func (server *Server) deleteRole(c *bytego.Ctx) error {
	var req deleteRoleRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	// 获取角色绑定的用户
	roles, err := server.store.GetRoleKeyByIDs(c.Context(), []int64{req.ID})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	groups := permission.Enforcer.GetFilteredNamedGroupingPolicy("g", 1, roles...)
	if len(groups) > 0 {
		return resp.Fail(http.StatusFound, code.RoleHasUserError).JSON(c)
	}

	// 获取角色绑定的菜单
	countRoleMenu, err := server.store.CountRoleMenuByRole(c.Context(), []int64{req.ID})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if countRoleMenu > 0 {
		return resp.Fail(http.StatusFound, code.RoleHasMenuError).JSON(c)
	}
	err = server.store.DeleteRole(c.Context(), []int64{req.ID})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.DelOK().JSON(c)
}

type batchDeleteRoleRequest struct {
	IDs []int64 `json:"ids" validate:"required"` // 主键集合
} // 批量删除角色请求参数

//@title 批量删除角色接口
//@api delete /sys/role/batch
//@group role
//@request batchDeleteRoleRequest
//@response 200 resp.resultOK{businesscode=10000,message="删除成功"}
func (server *Server) batchDeleteRole(c *bytego.Ctx) error {
	var req batchDeleteRoleRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	// 获取角色绑定的用户
	roles, err := server.store.GetRoleKeyByIDs(c.Context(), req.IDs)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	groups := permission.Enforcer.GetFilteredNamedGroupingPolicy("g", 1, roles...)
	if len(groups) > 0 {
		return resp.Fail(http.StatusFound, code.RoleHasUserError).JSON(c)
	}

	// 获取角色绑定的菜单
	countRoleMenu, err := server.store.CountRoleMenuByRole(c.Context(), req.IDs)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if countRoleMenu > 0 {
		return resp.Fail(http.StatusFound, code.RoleHasMenuError).JSON(c)
	}
	err = server.store.DeleteRole(c.Context(), req.IDs)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.DelOK().JSON(c)
}

type updateRolePermissionRequest struct {
	ID        int64                     `param:"id" validate:"required"` // 角色ID
	RoleMenus []db.CreateRoleMenuParams `json:"role_menus"`              // 角色菜单集合
} // 角色授权/解除菜单权限请求参数

//@title 角色授权/解除菜单权限接口
//@api post /sys/role/permission/:id
//@group role
//@request updateRolePermissionRequest
//@response 200 resp.resultOK{businesscode=10000,message="操作成功"}
func (server *Server) updateRolePermission(c *bytego.Ctx) error {
	var (
		req            updateRolePermissionRequest
		createRoleMenu []db.CreateRoleMenuParams
		errs           []error
	)
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	// 查询角色对应的菜单列表
	rm, err := server.store.ListRoleMenuByRole(c.Context(), req.ID)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
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
					return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(v).JSON(c)
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
				return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
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
				return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(v).JSON(c)
			}
		}
	}
	return resp.OperateOK().JSON(c)
}

type getRolePermissionRequest struct {
	ID int64 `param:"id" validate:"required"` // 角色ID
} // 获取角色授权的菜单权限请求参数

type getRolePermissionResponse struct {
	Menu   []int64           `json:"menu"`
	Button map[int64][]int64 `json:"button"`
} //获取角色授权的菜单权限

//@title 获取角色授权的菜单权限接口
//@api get /sys/role/permission/:id
//@group role
//@request getRolePermissionRequest
//@response 200 resp.resultOK{businesscode=10000,message="获取成功",data=getRolePermissionResponse}
func (server *Server) getRolePermission(c *bytego.Ctx) error {
	var req getRolePermissionRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	// 查询所有目录ID
	parentMenuIDs, err := server.store.ListMenuForParent(c.Context())
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	// 查询当前角色的菜单
	arg := db.ListRoleMenuForMenuParams{
		Role:         req.ID,
		Excludemenus: parentMenuIDs,
	}
	roleMenus, err := server.store.ListRoleMenuForMenu(c.Context(), arg)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	// 查询当前角色的按钮
	buttons, err := server.store.ListRoleMenuForButton(c.Context(), req.ID)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	menuButtons, err := server.store.ListMenuForParentIDByID(c.Context(), buttons)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	roleButtons := make(map[int64][]int64)
	for _, v := range menuButtons {
		roleButtons[v.Parent] = append(roleButtons[v.Parent], v.ID)
	}

	return resp.GetOK(getRolePermissionResponse{
		Menu:   roleMenus,
		Button: roleButtons,
	}).JSON(c)
}

type roleApiPermissionRequest struct {
	ID   int64   `param:"id" validate:"required"`            // 角色ID
	Type *int    `json:"type" validate:"required,oneof=0 1"` // 操作类型 0:解除api权限 1:绑定api权限
	Api  []int64 `json:"api" validate:"required"`            // 接口ID集合
} // 角色授权/解除接口权限请求参数

//@title 角色授权/解除菜单权限接口
//@api post /sys/role/api/:id
//@group role
//@request roleApiPermissionRequest
//@response 200 resp.resultOK{businesscode=10000,message="操作成功"}
func (server *Server) roleApiPermission(c *bytego.Ctx) error {
	var req roleApiPermissionRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	role, err := server.store.ListRoleByID(c.Context(), req.ID)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	apis, err := server.store.ListApiByIDs(c.Context(), req.Api)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	var groups [][]string
	for _, api := range apis {
		groups = append(groups, []string{role.Key, api.Url, api.Method})
	}

	if *req.Type == 1 {
		_, err = permission.Enforcer.AddNamedPolicies("p", groups)
		if err != nil {
			return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
		}
	} else if *req.Type == 0 {
		_, err := permission.Enforcer.RemoveNamedPolicies("p", groups)
		if err != nil {
			return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
		}
	}
	return resp.OperateOK().JSON(c)
}

type getRoleApiRequest struct {
	ID   int64 `param:"id" validate:"required"`  // 角色ID
	Menu int64 `json:"menu" validate:"required"` // 菜单ID
} //

//@title 获取角色授权的接口权限接口
//@api get /sys/role/api/:id
//@group role
//@request getRoleApiRequest
//@response 200 resp.resultOK{businesscode=10000,message="获取成功",data=[]int64} "接口ID集合"
func (server *Server) getRoleApi(c *bytego.Ctx) error {
	var (
		req  getRoleApiRequest
		errs []error
		ids  []int64
		apis []int64
	)
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	role, err := server.store.ListRoleByID(c.Context(), req.ID)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	roleApis := permission.Enforcer.GetFilteredNamedPolicy("p", 0, role.Key)
	if len(roleApis) > 0 {
		menuApiIds, err := server.store.ListMenuApiForApiByMenu(c.Context(), req.Menu)
		if err != nil {
			return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
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
	return resp.GetOK(apis).JSON(c)
}
