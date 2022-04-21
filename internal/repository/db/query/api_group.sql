-- name: ListApiGroup :many
SELECT * FROM api_groups
WHERE CASE WHEN @key::text = '' then 1=1 else name like concat('%',@key::text,'%') or remark like concat('%',@key::text,'%') end;

-- name: CreateApiGroup :exec
-- CreateApiGroup 创建 api组
INSERT INTO api_groups (
    name,remark
) VALUES (
    $1, $2
);