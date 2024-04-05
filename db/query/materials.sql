-- name: CreateMaterial :one
INSERT INTO materials (
    course_id,
    title,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetMaterial :one
SELECT * FROM materials
WHERE material_id = $1;

-- name: UpdateMaterial :one
UPDATE materials
SET title = $3, description = $4
WHERE material_id = $1 AND course_id = $2
RETURNING *;

-- name: ListMaterial :many
SELECT * FROM materials
WHERE course_id = $1
ORDER BY material_id
LIMIT $2
OFFSET $3;

-- name: DeleteMaterial :exec
DELETE FROM materials
WHERE material_id = $1 AND course_id = $2;