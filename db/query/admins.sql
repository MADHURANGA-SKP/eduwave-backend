-- name: GetAdmin :one
SELECT * FROM admins
WHERE admin_id = $1;

-- name: UpdateAdmin :one
UPDATE admins
SET user_name = $2
WHERE admin_id = $1
RETURNING *;

-- name: DeleteAdmin :exec
DELETE FROM admins
WHERE admin_id = $1;

