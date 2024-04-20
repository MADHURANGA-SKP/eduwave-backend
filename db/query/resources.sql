-- name: CreateResource :one
INSERT INTO resources (
    material_id,
    title,
    type,
    content_url
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetResource :one
SELECT * FROM resources
WHERE resource_id = $1;

-- name: UpdateResource :one
UPDATE resources
SET title = $2, type = $3, content_url = $4
WHERE resource_id = $1
RETURNING *;

-- name: ListResource :many
SELECT * FROM resources
ORDER BY resource_id
LIMIT $1
OFFSET $2;

-- name: ListResourceByMaterial :many
SELECT * FROM resources
WHERE material_id = $1
ORDER BY resource_id
LIMIT $2
OFFSET $3;

-- name: DeleteResource :exec
DELETE FROM resources
WHERE resource_id = $1;