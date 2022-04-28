// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: role.sql

package db

import (
	"context"
)

const createRole = `-- name: CreateRole :exec
INSERT INTO roles (
    name, is_disable, key, sort, remark
) VALUES (
    $1, $2, $3, $4, $5
)
`

type CreateRoleParams struct {
	Name      string  `json:"name"`
	IsDisable bool    `json:"is_disable"`
	Key       string  `json:"key"`
	Sort      int32   `json:"sort"`
	Remark    *string `json:"remark"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) error {
	_, err := q.db.Exec(ctx, createRole,
		arg.Name,
		arg.IsDisable,
		arg.Key,
		arg.Sort,
		arg.Remark,
	)
	return err
}

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = ANY($1::bigserial[])
`

func (q *Queries) DeleteRole(ctx context.Context, id []int64) error {
	_, err := q.db.Exec(ctx, deleteRole, id)
	return err
}

const getRoleKeyByIDs = `-- name: GetRoleKeyByIDs :many
SELECT key FROM roles
WHERE id = ANY($1::bigserial[])
`

func (q *Queries) GetRoleKeyByIDs(ctx context.Context, dollar_1 []int64) ([]string, error) {
	rows, err := q.db.Query(ctx, getRoleKeyByIDs, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var key string
		if err := rows.Scan(&key); err != nil {
			return nil, err
		}
		items = append(items, key)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRole = `-- name: ListRole :many
SELECT id, name, is_disable, key, sort, remark, created_at FROM roles
WHERE CASE WHEN $1::text = '' THEN 1=1 ELSE name like concat('%',$1::text,'%') END
AND CASE WHEN $2::text = '' THEN 1=1 ELSE key like concat('%',$2::text,'%') END
LIMIT $4::int
OFFSET $3::int
`

type ListRoleParams struct {
	Name       string `json:"name"`
	Key        string `json:"key"`
	Pageoffset int32  `json:"pageoffset"`
	Pagelimit  int32  `json:"pagelimit"`
}

func (q *Queries) ListRole(ctx context.Context, arg ListRoleParams) ([]*Role, error) {
	rows, err := q.db.Query(ctx, listRole,
		arg.Name,
		arg.Key,
		arg.Pageoffset,
		arg.Pagelimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Role{}
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.IsDisable,
			&i.Key,
			&i.Sort,
			&i.Remark,
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

const listRoleByID = `-- name: ListRoleByID :one
SELECT id, name, is_disable, key, sort, remark, created_at FROM roles
WHERE id = $1
`

func (q *Queries) ListRoleByID(ctx context.Context, id int64) (*Role, error) {
	row := q.db.QueryRow(ctx, listRoleByID, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsDisable,
		&i.Key,
		&i.Sort,
		&i.Remark,
		&i.CreatedAt,
	)
	return &i, err
}

const listRoleForIDByKeys = `-- name: ListRoleForIDByKeys :many
SELECT id FROM roles
WHERE key = ANY($1::text[])
`

func (q *Queries) ListRoleForIDByKeys(ctx context.Context, keys []string) ([]int64, error) {
	rows, err := q.db.Query(ctx, listRoleForIDByKeys, keys)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRole = `-- name: UpdateRole :exec
UPDATE roles
SET name = $1, is_disable = $2, key = $3, sort = $4, remark = $5
WHERE id = $6
`

type UpdateRoleParams struct {
	Name      string  `json:"name"`
	IsDisable bool    `json:"is_disable"`
	Key       string  `json:"key"`
	Sort      int32   `json:"sort"`
	Remark    *string `json:"remark"`
	ID        int64   `json:"id"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.db.Exec(ctx, updateRole,
		arg.Name,
		arg.IsDisable,
		arg.Key,
		arg.Sort,
		arg.Remark,
		arg.ID,
	)
	return err
}