-- name: ListApiGroup :many
SELECT * FROM api_groups
WHERE CASE WHEN @key::text = '' then 1=1 else name like concat('%',@key::text,'%') or remark like concat('%',@key::text,'%') end 
ORDER BY id
LIMIT @pageLimit::int 
OFFSET @pageOffset::int;

-- name: CreateApiGroup :exec
-- CreateApiGroup 创建 api 组
INSERT INTO api_groups (
    name,remark
) VALUES (
    $1, $2
);

-- name: UpdateApiGroup :exec
-- UpdateApiGroup 修改 api 组
UPDATE api_groups
SET name = $1, remark = $2
WHERE id = @id;

-- name: DeleteApiGroup :exec
-- DeleteApiGroup 删除 Api 组
DELETE FROM api_groups
WHERE id = ANY($1::bigint[]);

-- name: GetGroupByID :one
SELECT * FROM api_groups
WHERE id = $1;

