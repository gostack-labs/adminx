-- name: CountRoleMenuByRole :one
SELECT COUNT(*) FROM role_menus
WHERE role = ANY($1::bigint[]);


-- name: ListRoleMenuByRole :many
SELECT * FROM role_menus
WHERE role = $1;

-- name: CreateRoleMenu :batchexec
INSERT INTO role_menus (
    role, menu, type
) VALUES (
    $1, $2, $3
);

-- name: DeleteRoleMenu :exec
DELETE FROM role_menus
WHERE id = ANY($1::bigint[]);

-- name: ListRoleMenuForMenu :many
SELECT menu FROM role_menus
WHERE role = $1 and menu <> ANY(@excludeMenus::bigint[]);

-- name: ListRoleMenuForButton :many
SELECT menu FROM role_menus
WHERE role = $1 and type = 2;

-- name: ListRoleMenuForMenuByRoles :many
SELECT menu from role_menus
WHERE role = ANY(@roles::bigint[]) AND type = @type;
