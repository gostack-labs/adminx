-- name: ListMenusByType :many
SELECT * FROM menus
where type = ANY(@types::int[]);

-- name: CreateMenu :one
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
) RETURNING *;

-- name: UpdateMenu :one
UPDATE menus
SET title = $1, path = $2, name = $3, component = $4,
redirect = $5, hyperlink = $6, is_hide = $7, is_keep_alive = $8,
is_affix = $9, is_iframe = $10, auth = $11, icon = $12,
sort = $13
WHERE id = $14
RETURNING *;

-- name: DeleteMenu :exec
DELETE FROM menus WHERE id = ANY(@ids::bigint[]);

-- name: CountMenusByParent :one
SELECT count(*) FROM menus
WHERE parent = ANY(@parents::bigint[]);

-- name: ListMenuByParent :many
SELECT * FROM menus
WHERE parent = $1;

-- name: ListMenuForParent :many
-- ListMenuForParent 查询所有的目录
SELECT distinct parent FROM menus
WHERE parent != 0 and type = 2;

-- name: ListMenuForParentIDByID :many
SELECT id,parent FROM menus
WHERE id = ANY(@ids::bigserial[]);

-- name: ListMenuForAuthByIDs :many
SELECT DISTINCT UNNEST(auth) from menus
WHERE id = ANY(@ids::bigserial[]);

-- name: GetMenuByID :one
SELECT * FROM menus
WHERE id = $1;