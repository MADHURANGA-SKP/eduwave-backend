-- name: CreateAdmin :one
INSERT INTO admins(
    user_name
) VALUES (
    $1
) RETURNING *;

-- name: GetAdmin :one
SELECT * FROM admins
WHERE admin_id = $1;

-- name: UpdateAdmin :one
UPDATE admins
SET user_name = $2
WHERE admin_id = $1
RETURNING *;

-- name: ListAdmin :many
SELECT * FROM admins
WHERE user_name = $1
ORDER BY admin_id
LIMIT $2
OFFSET $3;

-- name: DeleteAdmin :exec
DELETE FROM admins
WHERE admin_id = $1;

