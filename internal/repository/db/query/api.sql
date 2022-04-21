-- name: ListApiByIDs :many
SELECT * FROM apis
WHERE id = ANY($1::bigserial[]);