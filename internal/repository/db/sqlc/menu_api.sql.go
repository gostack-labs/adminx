// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: menu_api.sql

package db

import (
	"context"
)

const deleteMenuApiByMenuAndApi = `-- name: DeleteMenuApiByMenuAndApi :exec
DELETE FROM menu_apis
WHERE menu = $1 AND api = ANY($2::bigint[])
`

type DeleteMenuApiByMenuAndApiParams struct {
	Menu int64   `json:"menu"`
	Apis []int64 `json:"apis"`
}

func (q *Queries) DeleteMenuApiByMenuAndApi(ctx context.Context, arg DeleteMenuApiByMenuAndApiParams) error {
	_, err := q.db.Exec(ctx, deleteMenuApiByMenuAndApi, arg.Menu, arg.Apis)
	return err
}

const listMenuApiByApi = `-- name: ListMenuApiByApi :many
SELECT id, menu, api, created_at FROM menu_apis
WHERE api = ANY($1::bigint[])
`

func (q *Queries) ListMenuApiByApi(ctx context.Context, api []int64) ([]*MenuApi, error) {
	rows, err := q.db.Query(ctx, listMenuApiByApi, api)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*MenuApi{}
	for rows.Next() {
		var i MenuApi
		if err := rows.Scan(
			&i.ID,
			&i.Menu,
			&i.Api,
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

const listMenuApiForApiByMenu = `-- name: ListMenuApiForApiByMenu :many
SELECT api FROM menu_apis
WHERE menu = $1
`

func (q *Queries) ListMenuApiForApiByMenu(ctx context.Context, menu int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, listMenuApiForApiByMenu, menu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var api int64
		if err := rows.Scan(&api); err != nil {
			return nil, err
		}
		items = append(items, api)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
