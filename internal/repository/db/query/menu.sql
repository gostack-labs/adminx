-- name: ListMenusByType :many
SELECT * FROM menus
where type = ANY($1::int[]);

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

-- name: DeleteMenu :exec
DELETE FROM menus WHERE id = ANY($1::bigserial[]);

-- name: CountMenusByParent :one
SELECT count(*) FROM menus
WHERE parent = ANY($1::bigserial[]);

-- name: ListMenuByParent :many
SELECT * FROM menus
WHERE parent = $1;