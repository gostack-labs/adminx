-- name: ListApiByIDs :many
SELECT * FROM apis
WHERE id = ANY($1::bigint[]);

-- name: ListApiByGroup :many
SELECT * FROM apis
WHERE groups = ANY($1::bigint[]);