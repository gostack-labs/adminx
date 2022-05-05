-- name: ListRole :many
SELECT * FROM roles
WHERE CASE WHEN @name::text = '' THEN 1=1 ELSE name like concat('%',@name::text,'%') END
AND CASE WHEN @key::text = '' THEN 1=1 ELSE key like concat('%',@key::text,'%') END
LIMIT @pageLimit::int
OFFSET @pageOffset::int;

-- name: GetRoleKeyByIDs :many
SELECT key FROM roles
WHERE id = ANY($1::bigint[]);

-- name: CreateRole :exec
INSERT INTO roles (
    name, is_disable, key, sort, remark
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: UpdateRole :exec
UPDATE roles
SET name = $1, is_disable = $2, key = $3, sort = $4, remark = $5
WHERE id = $6;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = ANY(@id::bigint[]);

-- name: ListRoleByID :one
SELECT * FROM roles
WHERE id = $1;

-- name: ListRoleForIDByKeys :many
SELECT id FROM roles
WHERE key = ANY(@keys::text[]);
