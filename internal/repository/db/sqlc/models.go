// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// 接口表
type Api struct {
	ID int64 `json:"id"`
	// 标题
	Title string `json:"title"`
	// 接口地址
	Url string `json:"url"`
	// 请求方式
	Method string `json:"method"`
	// 分组
	Group int64 `json:"group"`
	// 备注
	Remark    sql.NullString `json:"remark"`
	CreatedAt time.Time      `json:"created_at"`
}

// 接口组表
type ApiGroup struct {
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 备注
	Remark    sql.NullString `json:"remark"`
	CreatedAt time.Time      `json:"created_at"`
}

type CasbinRule struct {
	ID        int64          `json:"id"`
	PType     string         `json:"p_type"`
	V0        sql.NullString `json:"v0"`
	V1        sql.NullString `json:"v1"`
	V2        sql.NullString `json:"v2"`
	V3        sql.NullString `json:"v3"`
	V4        sql.NullString `json:"v4"`
	V5        sql.NullString `json:"v5"`
	CreatedAt time.Time      `json:"created_at"`
}

// 菜单表
type Menu struct {
	ID int64 `json:"id"`
	// 父级
	Parent int64 `json:"parent"`
	// 标题
	Title string `json:"title"`
	// 路径
	Path sql.NullString `json:"path"`
	// 路由名称
	Name string `json:"name"`
	// 组件路径
	Component sql.NullString `json:"component"`
	// 跳转路径
	Redirect sql.NullString `json:"redirect"`
	// 超链接
	Hyperlink sql.NullString `json:"hyperlink"`
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
	Icon sql.NullString `json:"icon"`
	// 类型：1 目录，2 菜单，3 按钮
	Type int32 `json:"type"`
	// 顺序
	Sort      int32     `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
}

// 菜单接口关联表
type MenuApi struct {
	ID int64 `json:"id"`
	// 菜单
	Menu int64 `json:"menu"`
	// 接口
	Api       int64     `json:"api"`
	CreatedAt time.Time `json:"created_at"`
}

// 角色表
type Role struct {
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 是否禁用
	IsDisable bool `json:"is_disable"`
	// 标识
	Key string `json:"key"`
	// 排序
	Sort int32 `json:"sort"`
	// 备注
	Remark    sql.NullString `json:"remark"`
	CreatedAt time.Time      `json:"created_at"`
}

// 角色菜单关联表
type RoleMenu struct {
	ID int64 `json:"id"`
	// 角色
	Role int64 `json:"role"`
	// 菜单
	Menu int64 `json:"menu"`
	// 类型：1 菜单，2 按钮
	Type      int32     `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Session struct {
	ID uuid.UUID `json:"id"`
	// 用户名，关联Users表username字段
	Username string `json:"username"`
	// 刷新密钥
	RefreshToken string `json:"refresh_token"`
	// 用户代理
	UserAgent string `json:"user_agent"`
	// ip
	ClientIp string `json:"client_ip"`
	// 是否屏蔽
	IsBlocked bool `json:"is_blocked"`
	// 过期时间
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	// 主键，用户名
	Username string `json:"username"`
	// 加密后密码
	HashedPassword string `json:"hashed_password"`
	// 全名
	FullName string `json:"full_name"`
	// 邮箱
	Email string `json:"email"`
	// 手机号
	Phone string `json:"phone"`
	// 修改密码时间
	PasswordChangeAt time.Time `json:"password_change_at"`
	CreatedAt        time.Time `json:"created_at"`
}