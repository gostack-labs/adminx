package api

import (
	"net/http"

	"github.com/gostack-labs/adminx/internal/code"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgx/v4"
)

type MenuValue struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	Path      *string `json:"path"`
	Name      string  `json:"name"`
	Component *string `json:"component"`
	Parent    int64   `json:"parent"`
	Type      int32   `json:"type"`
	Sort      int32   `json:"sort"`
	Meta      struct {
		Title       string   `json:"title"`
		Hyperlink   *string  `json:"hyperlink"`
		IsHide      bool     `json:"is_hide"`
		IsKeepAlive bool     `json:"is_keep_alive"`
		IsAffix     bool     `json:"is_affix"`
		IsIframe    bool     `json:"is_iframe"`
		Auth        []string `json:"auth"`
		Icon        *string  `json:"icon"`
	} `json:"meta"`
	Children []*MenuValue `json:"children"`
}

type MenuTree struct {
	Menus       []*db.Menu
	ParentMenus map[int64][]*MenuValue
}

func (m *MenuTree) formatMenus(menuValueList []*MenuValue) {
	m.ParentMenus = make(map[int64][]*MenuValue)
	for _, menu := range menuValueList {
		if menu.Parent != 0 {
			m.ParentMenus[menu.Parent] = append(m.ParentMenus[menu.Parent], menu)
		}
	}
}

func (m *MenuTree) recursionMenuTree(menus []*MenuValue) {
	for _, menu := range menus {
		if _, ok := m.ParentMenus[menu.ID]; ok {
			menu.Children = m.ParentMenus[menu.ID]
			m.recursionMenuTree(menu.Children)
		}
	}
}

func (m *MenuTree) GetMenuTree() []*MenuValue {
	var (
		topMenuList   []*MenuValue
		menuValueList []*MenuValue
	)

	topMenuList = make([]*MenuValue, 0)
	for _, menu := range m.Menus {
		menuValue := &MenuValue{
			ID:        menu.ID,
			Title:     menu.Title,
			Path:      menu.Path,
			Name:      menu.Name,
			Component: menu.Component,
			Parent:    menu.Parent,
			Type:      menu.Type,
			Sort:      menu.Sort,
			Meta: struct {
				Title       string   `json:"title"`
				Hyperlink   *string  `json:"hyperlink"`
				IsHide      bool     `json:"is_hide"`
				IsKeepAlive bool     `json:"is_keep_alive"`
				IsAffix     bool     `json:"is_affix"`
				IsIframe    bool     `json:"is_iframe"`
				Auth        []string `json:"auth"`
				Icon        *string  `json:"icon"`
			}{
				Title:       menu.Title,
				Hyperlink:   menu.Hyperlink,
				IsHide:      menu.IsHide,
				IsKeepAlive: menu.IsKeepAlive,
				IsAffix:     menu.IsAffix,
				IsIframe:    menu.IsIframe,
				Auth:        menu.Auth,
				Icon:        menu.Icon,
			},
			Children: []*MenuValue{},
		}
		menuValueList = append(menuValueList, menuValue)
		if menu.Parent == 0 {
			topMenuList = append(topMenuList, menuValue)
		}
	}

	m.formatMenus(menuValueList)

	m.recursionMenuTree(topMenuList)
	return topMenuList
}

func (server *Server) menuTree(c *bytego.Ctx) error {
	var (
		err        error
		menus      []*db.Menu
		result     []*MenuValue
		m          MenuTree
		buttonList []*db.Menu
	)
	menus, err = server.store.ListMenusByType(c.Context(), []int32{1, 2})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	m = MenuTree{Menus: menus}
	result = m.GetMenuTree()

	buttonList, err = server.store.ListMenusByType(c.Context(), []int32{3})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	buttonMap := make(map[int64][]*db.Menu)
	for _, b := range buttonList {
		buttonMap[b.Parent] = append(buttonMap[b.Parent], b)
	}
	return resp.GetOK(bytego.Map{
		"menu":   result,
		"button": buttonMap,
	}).JSON(c)
}

type createMenuRequest struct {
	// 父级
	Parent int64 `json:"parent" validate:"required,numeric"`
	// 标题
	Title string `json:"title" validate:"required"`
	// 路径
	Path *string `json:"path"`
	// 路由名称
	Name string `json:"name" validate:"required"`
	// 组件路径
	Component *string `json:"component"`
	// 跳转路径
	Redirect *string `json:"redirect"`
	// 超链接
	Hyperlink *string `json:"hyperlink"`
	// 是否隐藏
	IsHide bool `json:"is_hide"`
	// 是否缓存组件状态
	IsKeepAlive bool `json:"is_keep_alive"`
	// 是否固定在标签栏
	IsAffix bool `json:"is_affix"`
	// 是否内嵌窗口
	IsIframe bool `json:"is_iframe"`
	// 权限粒子
	Auth []string `json:"auth"`
	// 图标
	Icon *string `json:"icon"`
	// 类型：1 目录，2 菜单，3 按钮
	Type int32 `json:"type" validate:"oneof=1 2 3"`
	// 顺序
	Sort int32 `json:"sort"`
}

func (server *Server) createMenu(c *bytego.Ctx) error {
	var req createMenuRequest

	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	arg := db.CreateMenuParams{
		Parent:      req.Parent,
		Title:       req.Title,
		Path:        req.Path,
		Name:        req.Name,
		Component:   req.Component,
		Redirect:    req.Redirect,
		Hyperlink:   req.Hyperlink,
		IsHide:      req.IsHide,
		IsKeepAlive: req.IsKeepAlive,
		IsAffix:     req.IsAffix,
		IsIframe:    req.IsIframe,
		Auth:        req.Auth,
		Icon:        req.Icon,
		Type:        req.Type,
		Sort:        req.Sort,
	}
	if req.Parent == 0 {
		if req.Type == 3 {
			return resp.Fail(http.StatusBadRequest, code.MenuRootTypeError).JSON(c)
		} else {
			menu, err := server.store.CreateMenu(c.Context(), arg)
			if err != nil {
				return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
			}
			return resp.CreateOK(menu).JSON(c)
		}
	}
	m, err := server.store.GetMenuByID(c.Context(), req.Parent)
	if err != nil {
		if err == pgx.ErrNoRows {
			return resp.Fail(http.StatusNotFound, code.MenuParentNoRowError).JSON(c)
		}
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	bussinessCode := 0
	switch m.Type {
	case 1:
		if req.Type == 3 {
			bussinessCode = code.MenuNoButtonError
		}
	case 2:
		if req.Type == 1 || req.Type == 2 {
			bussinessCode = code.MenuNoDirOrMenuError
		}
	case 3:
		bussinessCode = code.MenuNoChildNodeError
	}
	if bussinessCode > 0 {
		return resp.Fail(http.StatusBadRequest, bussinessCode).JSON(c)
	}

	menu, err := server.store.CreateMenu(c.Context(), arg)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.CreateOK(menu).JSON(c)
}

type updateMenuRequest struct {
	ID    int64  `param:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
	// 路径
	Path *string `json:"path"`
	// 路由名称
	Name string `json:"name" validate:"required"`
	// 组件路径
	Component *string `json:"component"`
	// 跳转路径
	Redirect *string `json:"redirect"`
	// 超链接
	Hyperlink *string `json:"hyperlink"`
	// 是否隐藏
	IsHide bool `json:"is_hide"`
	// 是否缓存组件状态
	IsKeepAlive bool `json:"is_keep_alive"`
	// 是否固定在标签栏
	IsAffix bool `json:"is_affix"`
	// 是否内嵌窗口
	IsIframe bool `json:"is_iframe"`
	// 权限粒子
	Auth []string `json:"auth"`
	// 图标
	Icon *string `json:"icon"`
	// 顺序
	Sort int32 `json:"sort"`
}

func (server *Server) updateMenu(c *bytego.Ctx) error {
	var req updateMenuRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	arg := db.UpdateMenuParams{
		ID:          req.ID,
		Title:       req.Title,
		Path:        req.Path,
		Name:        req.Name,
		Component:   req.Component,
		Redirect:    req.Redirect,
		Hyperlink:   req.Hyperlink,
		IsHide:      req.IsHide,
		IsKeepAlive: req.IsKeepAlive,
		IsAffix:     req.IsAffix,
		IsIframe:    req.IsIframe,
		Auth:        req.Auth,
		Icon:        req.Icon,
		Sort:        req.Sort,
	}
	menu, err := server.store.UpdateMenu(c.Context(), arg)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.UpdateOK(menu).JSON(c)
}

type deleteMenuRequest struct {
	ID int64 `param:"id" validate:"required,numeric"`
}

func (server *Server) deleteMenu(c *bytego.Ctx) error {
	var req deleteMenuRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	// Check whether child nodes exist. If yes, they cannot be deleted
	menuCount, err := server.store.CountMenusByParent(c.Context(), []int64{req.ID})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if menuCount > 0 {
		return resp.Fail(http.StatusFound, code.MenuHasChildNodeError).JSON(c)
	}
	err = server.store.DeleteMenu(c.Context(), []int64{req.ID})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.DelOK().JSON(c)
}

type batchDeleteMenuRequest struct {
	IDs []int64 `json:"ids" validate:"required"`
}

func (server *Server) batchDeleteMenu(c *bytego.Ctx) error {
	var req batchDeleteMenuRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	menuCount, err := server.store.CountMenusByParent(c.Context(), req.IDs)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	if menuCount > 0 {
		return resp.Fail(http.StatusFound, code.MenuHasChildNodeError).JSON(c)
	}
	err = server.store.DeleteMenu(c.Context(), req.IDs)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).JSON(c)
	}
	return resp.DelOK().JSON(c)
}

// 查询菜单下的所有按钮
type menuButtonRequest struct {
	ID int64 `param:"id" validate:"required,numeric"`
}

func (server *Server) menuButton(c *bytego.Ctx) error {
	var req menuButtonRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	buttonList, err := server.store.ListMenuByParent(c.Context(), req.ID)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.GetOK(buttonList).JSON(c)
}

type menuBindApiRequest struct {
	ID   int64   `param:"id" validate:"required"`
	Type int     `json:"type" validate:"required,oneof=1 2"` // 1:bind 2:unbind
	Apis []int64 `json:"apis" validate:"required"`
}

func (server *Server) mentBindApi(c *bytego.Ctx) error {
	var (
		req  menuBindApiRequest
		errs []error
	)

	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	if req.Type == 1 {
		existApis, err := server.store.ListMenuApiForApiByMenu(c.Context(), req.ID)
		if err != nil {
			return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
		}
		apiMaps := make(map[int64]struct{})
		if len(existApis) != 0 {
			for _, i := range existApis {
				apiMaps[i] = struct{}{}
			}
		}

		args := make([]db.CreateMenuApiParams, 0)
		for _, i := range req.Apis {
			if _, ok := apiMaps[i]; !ok {
				args = append(args, db.CreateMenuApiParams{
					Menu: req.ID,
					Api:  i,
				})
			}
		}
		if len(args) > 0 {
			cma := server.store.CreateMenuApi(c.Context(), args)
			defer cma.Close()
			cma.Exec(func(i int, err error) {
				errs[i] = err
			})
			for _, v := range errs {
				if v != nil {
					return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(v).JSON(c)
				}
			}
		}
	} else if req.Type == 0 {
		arg := db.DeleteMenuApiByMenuAndApiParams{
			Menu: req.ID,
			Apis: req.Apis,
		}
		err := server.store.DeleteMenuApiByMenuAndApi(c.Context(), arg)
		if err != nil {
			return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
		}
	}
	return resp.OperateOK().JSON(c)
}

type mentApisRequest struct {
	Menu int64 `param:"menu" validata:"required"`
}

func (server *Server) MenuApis(c *bytego.Ctx) error {
	var req mentApisRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	apiList, err := server.store.ListMenuApiForApiByMenu(c.Context(), req.Menu)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.GetOK(apiList).JSON(c)
}

type menuApiListRequest struct {
	Menu int64 `param:"menu" validate:"required"`
}

func (server *Server) MenuApiList(c *bytego.Ctx) error {
	var req menuApiListRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	apiIDs, err := server.store.ListMenuApiForApiByMenu(c.Context(), req.Menu)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	apiList, err := server.store.ListApiByIDs(c.Context(), apiIDs)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	return resp.GetOK(apiList).JSON(c)
}
