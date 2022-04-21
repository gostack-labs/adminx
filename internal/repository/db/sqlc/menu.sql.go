// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: menu.sql

package db

import (
	"context"
)

const countMenusByParent = `-- name: CountMenusByParent :one
SELECT count(*) FROM menus
WHERE parent = ANY($1::bigint[])
`

func (q *Queries) CountMenusByParent(ctx context.Context, dollar_1 []int64) (int64, error) {
	row := q.db.QueryRow(ctx, countMenusByParent, dollar_1)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createMenu = `-- name: CreateMenu :one
INSERT INTO menus (
    parent,
    title,
    path,
    name,
    component,
    redirect,
    hyperlink,
    is_hide,
    is_keep_alive,
    is_affix,
    is_iframe,
    auth,
    icon,
    type,
    sort
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
) RETURNING id, parent, title, path, name, component, redirect, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, type, sort, created_at
`

type CreateMenuParams struct {
	Parent      int64    `json:"parent"`
	Title       string   `json:"title"`
	Path        *string  `json:"path"`
	Name        string   `json:"name"`
	Component   *string  `json:"component"`
	Redirect    *string  `json:"redirect"`
	Hyperlink   *string  `json:"hyperlink"`
	IsHide      bool     `json:"is_hide"`
	IsKeepAlive bool     `json:"is_keep_alive"`
	IsAffix     bool     `json:"is_affix"`
	IsIframe    bool     `json:"is_iframe"`
	Auth        []string `json:"auth"`
	Icon        *string  `json:"icon"`
	Type        int32    `json:"type"`
	Sort        int32    `json:"sort"`
}

func (q *Queries) CreateMenu(ctx context.Context, arg CreateMenuParams) (*Menu, error) {
	row := q.db.QueryRow(ctx, createMenu,
		arg.Parent,
		arg.Title,
		arg.Path,
		arg.Name,
		arg.Component,
		arg.Redirect,
		arg.Hyperlink,
		arg.IsHide,
		arg.IsKeepAlive,
		arg.IsAffix,
		arg.IsIframe,
		arg.Auth,
		arg.Icon,
		arg.Type,
		arg.Sort,
	)
	var i Menu
	err := row.Scan(
		&i.ID,
		&i.Parent,
		&i.Title,
		&i.Path,
		&i.Name,
		&i.Component,
		&i.Redirect,
		&i.Hyperlink,
		&i.IsHide,
		&i.IsKeepAlive,
		&i.IsAffix,
		&i.IsIframe,
		&i.Auth,
		&i.Icon,
		&i.Type,
		&i.Sort,
		&i.CreatedAt,
	)
	return &i, err
}

const deleteMenu = `-- name: DeleteMenu :exec
DELETE FROM menus WHERE id = ANY($1::bigint[])
`

func (q *Queries) DeleteMenu(ctx context.Context, dollar_1 []int64) error {
	_, err := q.db.Exec(ctx, deleteMenu, dollar_1)
	return err
}

const listMenuByParent = `-- name: ListMenuByParent :many
SELECT id, parent, title, path, name, component, redirect, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, type, sort, created_at FROM menus
WHERE parent = $1
`

func (q *Queries) ListMenuByParent(ctx context.Context, parent int64) ([]*Menu, error) {
	rows, err := q.db.Query(ctx, listMenuByParent, parent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Menu{}
	for rows.Next() {
		var i Menu
		if err := rows.Scan(
			&i.ID,
			&i.Parent,
			&i.Title,
			&i.Path,
			&i.Name,
			&i.Component,
			&i.Redirect,
			&i.Hyperlink,
			&i.IsHide,
			&i.IsKeepAlive,
			&i.IsAffix,
			&i.IsIframe,
			&i.Auth,
			&i.Icon,
			&i.Type,
			&i.Sort,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMenusByType = `-- name: ListMenusByType :many
SELECT id, parent, title, path, name, component, redirect, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, type, sort, created_at FROM menus
where type = ANY($1::int[])
`

func (q *Queries) ListMenusByType(ctx context.Context, dollar_1 []int32) ([]*Menu, error) {
	rows, err := q.db.Query(ctx, listMenusByType, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Menu{}
	for rows.Next() {
		var i Menu
		if err := rows.Scan(
			&i.ID,
			&i.Parent,
			&i.Title,
			&i.Path,
			&i.Name,
			&i.Component,
			&i.Redirect,
			&i.Hyperlink,
			&i.IsHide,
			&i.IsKeepAlive,
			&i.IsAffix,
			&i.IsIframe,
			&i.Auth,
			&i.Icon,
			&i.Type,
			&i.Sort,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
