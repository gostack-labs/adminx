-- name: ListMenuApiForApiByMenu :many
SELECT api FROM menu_apis
WHERE menu = $1;

-- name: CreateMenuApi :batchexec
INSERT INTO menu_apis (
    menu,
    api
) VALUES (
    $1, $2
);

-- name: DeleteMenuApiByMenuAndApi :exec
DELETE FROM menu_apis
WHERE menu = $1 AND api = ANY(@apis::bigint[]);

-- name: ListMenuApiByApi :many
SELECT * FROM menu_apis
WHERE api = ANY(@api::bigint[]);