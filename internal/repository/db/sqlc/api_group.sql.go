// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: api_group.sql

package db

import (
	"context"
)

const createApiGroup = `-- name: CreateApiGroup :exec
INSERT INTO api_groups (
    name,remark
) VALUES (
    $1, $2
)
`

type CreateApiGroupParams struct {
	Name   string  `json:"name"`
	Remark *string `json:"remark"`
}

// CreateApiGroup 创建 api 组
func (q *Queries) CreateApiGroup(ctx context.Context, arg CreateApiGroupParams) error {
	_, err := q.db.Exec(ctx, createApiGroup, arg.Name, arg.Remark)
	return err
}

const deleteApiGroup = `-- name: DeleteApiGroup :exec
DELETE FROM api_groups
WHERE id = ANY($1::bigint[])
`

// DeleteApiGroup 删除 Api 组
func (q *Queries) DeleteApiGroup(ctx context.Context, dollar_1 []int64) error {
	_, err := q.db.Exec(ctx, deleteApiGroup, dollar_1)
	return err
}

const listApiGroup = `-- name: ListApiGroup :many
SELECT id, name, remark, created_at FROM api_groups
WHERE CASE WHEN $1::text = '' then 1=1 else name like concat('%',$1::text,'%') or remark like concat('%',$1::text,'%') end 
ORDER BY id
LIMIT $3::int 
OFFSET $2::int
`

type ListApiGroupParams struct {
	Key        string `json:"key"`
	Pageoffset int32  `json:"pageoffset"`
	Pagelimit  int32  `json:"pagelimit"`
}

func (q *Queries) ListApiGroup(ctx context.Context, arg ListApiGroupParams) ([]*ApiGroup, error) {
	rows, err := q.db.Query(ctx, listApiGroup, arg.Key, arg.Pageoffset, arg.Pagelimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ApiGroup{}
	for rows.Next() {
		var i ApiGroup
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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

const updateApiGroup = `-- name: UpdateApiGroup :exec
UPDATE api_groups
SET name = $1, remark = $2
WHERE id = $3
`

type UpdateApiGroupParams struct {
	Name   string  `json:"name"`
	Remark *string `json:"remark"`
	ID     int64   `json:"id"`
}

// UpdateApiGroup 修改 api 组
func (q *Queries) UpdateApiGroup(ctx context.Context, arg UpdateApiGroupParams) error {
	_, err := q.db.Exec(ctx, updateApiGroup, arg.Name, arg.Remark, arg.ID)
	return err
}
