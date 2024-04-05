-- name: GetRole :one
SELECT * FROM roles
WHERE role_id = $1 LIMIT 1;
