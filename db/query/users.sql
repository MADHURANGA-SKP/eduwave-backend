-- name: CreateUser :one
INSERT INTO users (
    user_name,
    full_name,
    hashed_password,
    email,
    role,
    qualification
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- -- name: GetUser :one
-- SELECT users.user_name AS user_username, teachers.user_name AS teacher_username, admins.user_name AS admin_username
-- FROM users
-- LEFT JOIN teachers ON users.user_id = teachers.user_id
-- LEFT JOIN admins ON users.user_id = admins.user_id
-- WHERE 
--     users.user_name = $1 OR teachers.user_name = $1 OR admins.user_name = $1;


-- name: GetUser :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;


-- name: GetUserById :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET 
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email),
    is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
    user_name = sqlc.arg(user_name)
RETURNING *;

-- name: ListUser :many
SELECT * FROM users
WHERE role = $1
ORDER BY user_id
LIMIT $2
OFFSET $3;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE user_id = $1;