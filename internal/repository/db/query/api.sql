-- name: ListApiByIDs :many
SELECT * FROM apis
WHERE id = ANY($1::bigint[]);

-- name: ListApiByGroup :many
SELECT * FROM apis
WHERE groups = ANY($1::bigint[]);

-- name: ListApi :many
SELECT * FROM apis
WHERE groups = @groups::bigint
AND CASE WHEN @title::text = '' THEN 1=1 ELSE title like concat('%',@title::text,'%') END
LIMIT @pageLimit::int
OFFSET @pageOffset::int;

-- name: CreateApi :exec
INSERT INTO apis (
    title, url, method, groups, remark
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: UpdateApi :exec
UPDATE apis
SET title = $1, url = $2, method = $3, groups = $4, remark = $5
WHERE id = $6;

-- name: DeleteApi :exec
DELETE FROM apis
WHERE id = ANY(@id::bigint[]);

-- name: ListApiBatch :batchmany
SELECT id FROM apis
WHERE url = $1 AND method = $2;

-- name: CountApiByMUT :one
-- CountApiByMUT 根据 标题 url method 查询数量
SELECT count(*) FROM apis
WHERE title = $1 AND url = $2 AND method = $3;
