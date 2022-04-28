-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email,
    phone
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE (username = $1 or email = $1 or phone = $1) LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = $1 LIMIT 1;


-- name: ListUser :many
SELECT * FROM users
WHERE CASE WHEN @username::text = '' THEN 1=1 ELSE name like concat('%',@username::text,'%') END
AND CASE WHEN @fullName::text = '' THEN 1=1 ELSE key like concat('%',@fullName::text,'%') END
AND CASE WHEN @email::text = '' THEN 1=1 ELSE key like concat('%',@email::text,'%') END
AND CASE WHEN @phone::text = '' THEN 1=1 ELSE key like concat('%',@phone::text,'%') END
LIMIT @pageLimit::int
OFFSET @pageOffset::int;

-- name: UpdateUser :exec
UPDATE users
SET full_name = $1, email = $2, phone = $3
WHERE username = @username::text;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;