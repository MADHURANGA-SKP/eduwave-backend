-- name: CreateAdmin :one
INSERT INTO admins (
    user_id,
    full_name,
    user_name,
    email,
    hashed_password
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;


-- name: GetAdmin :one
SELECT * FROM admins
WHERE admin_id = $1;

-- name: UpdateAdmin :one
UPDATE admins
SET full_name = $2, user_name = $3, email = $4, hashed_password = $5
WHERE admin_id = $1
RETURNING *;

-- name: DeleteAdmin :exec
DELETE FROM admins
WHERE admin_id = $1;

