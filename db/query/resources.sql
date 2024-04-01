-- name: CreateResource :one
INSERT INTO resources (
    title,
    type,
    content_url
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetResource :one
SELECT * FROM resources
WHERE material_id = $1 AND resource_id = $2;

-- name: UpdateResource :one
UPDATE resources
SET title = $3, type = $4, content_url = $5
WHERE material_id = $1 AND resource_id = $2
RETURNING *;

-- name: ListResource :many
SELECT * FROM resources
WHERE material_id = $1 AND resource_id = $2
ORDER BY resource_id
LIMIT $3
OFFSET $4;

-- name: DeleteResource :exec
DELETE FROM resources
WHERE resource_id = $1 AND material_id = $2;