-- name: CreateUser :one
INSERT INTO users (
    user_name,
    role,
    full_name,
    hashed_password,
    email
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING*;

-- name: GetUser :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;

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
