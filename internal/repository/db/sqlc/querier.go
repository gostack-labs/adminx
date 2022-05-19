// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CheckGroupExist(ctx context.Context, id int64) (bool, error)
	CheckUserEmail(ctx context.Context, email string) (bool, error)
	CheckUserPhone(ctx context.Context, phone string) (bool, error)
	// CountApiByMUT 根据 标题 url method 查询数量
	CountApiByMUT(ctx context.Context, arg CountApiByMUTParams) (int64, error)
	CountMenusByParent(ctx context.Context, parents []int64) (int64, error)
	CountRoleMenuByRole(ctx context.Context, dollar_1 []int64) (int64, error)
	CreateApi(ctx context.Context, arg CreateApiParams) error
	// CreateApiGroup 创建 api 组
	CreateApiGroup(ctx context.Context, arg CreateApiGroupParams) error
	CreateMenu(ctx context.Context, arg CreateMenuParams) (*Menu, error)
	CreateMenuApi(ctx context.Context, arg []CreateMenuApiParams) *CreateMenuApiBatchResults
	CreateRole(ctx context.Context, arg CreateRoleParams) error
	CreateRoleMenu(ctx context.Context, arg []CreateRoleMenuParams) *CreateRoleMenuBatchResults
	CreateSession(ctx context.Context, arg CreateSessionParams) (*Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)
	DeleteApi(ctx context.Context, id []int64) error
	// DeleteApiGroup 删除 Api 组
	DeleteApiGroup(ctx context.Context, dollar_1 []int64) error
	DeleteMenu(ctx context.Context, ids []int64) error
	DeleteMenuApiByMenuAndApi(ctx context.Context, arg DeleteMenuApiByMenuAndApiParams) error
	DeleteRole(ctx context.Context, id []int64) error
	DeleteRoleMenu(ctx context.Context, dollar_1 []int64) error
	DeleteUser(ctx context.Context, dollar_1 []string) error
	GetGroupByID(ctx context.Context, id int64) (*ApiGroup, error)
	GetMenuByID(ctx context.Context, id int64) (*Menu, error)
	GetRoleKeyByIDs(ctx context.Context, dollar_1 []int64) ([]string, error)
	GetSession(ctx context.Context, id uuid.UUID) (*Session, error)
	GetUser(ctx context.Context, username string) (*User, error)
	ListApi(ctx context.Context, arg ListApiParams) ([]*Api, error)
	ListApiBatch(ctx context.Context, arg []ListApiBatchParams) *ListApiBatchBatchResults
	ListApiByGroup(ctx context.Context, dollar_1 []int64) ([]*Api, error)
	ListApiByIDs(ctx context.Context, dollar_1 []int64) ([]*Api, error)
	ListApiGroup(ctx context.Context, arg ListApiGroupParams) ([]*ApiGroup, error)
	ListMenuApiByApi(ctx context.Context, api []int64) ([]*MenuApi, error)
	ListMenuApiForApiByMenu(ctx context.Context, menu int64) ([]int64, error)
	ListMenuByParent(ctx context.Context, parent int64) ([]*Menu, error)
	ListMenuForAuthByIDs(ctx context.Context, ids []int64) ([]interface{}, error)
	// ListMenuForParent 查询所有的目录
	ListMenuForParent(ctx context.Context) ([]int64, error)
	ListMenuForParentIDByID(ctx context.Context, ids []int64) ([]*ListMenuForParentIDByIDRow, error)
	ListMenusByType(ctx context.Context, types []int32) ([]*Menu, error)
	ListRole(ctx context.Context, arg ListRoleParams) ([]*Role, error)
	ListRoleByID(ctx context.Context, id int64) (*Role, error)
	ListRoleForIDByKeys(ctx context.Context, keys []string) ([]int64, error)
	ListRoleMenuByRole(ctx context.Context, role int64) ([]*RoleMenu, error)
	ListRoleMenuForButton(ctx context.Context, role int64) ([]int64, error)
	ListRoleMenuForMenu(ctx context.Context, arg ListRoleMenuForMenuParams) ([]int64, error)
	ListRoleMenuForMenuByRoles(ctx context.Context, arg ListRoleMenuForMenuByRolesParams) ([]int64, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]*User, error)
	UpdateApi(ctx context.Context, arg UpdateApiParams) error
	// UpdateApiGroup 修改 api 组
	UpdateApiGroup(ctx context.Context, arg UpdateApiGroupParams) error
	UpdateMenu(ctx context.Context, arg UpdateMenuParams) (*Menu, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
