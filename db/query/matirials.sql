-- name: CreateMatirials :one
INSERT INTO matirials (
    title,
    description
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetMatirials :one
SELECT * FROM matirials
WHERE course_id = $1;

-- name: UpdateMatirials :one
UPDATE matirials
SET title = $3, description = $4
WHERE matirial_id = $1 AND course_id = $2
RETURNING *;

-- name: ListMatirials :many
SELECT * FROM matirials
WHERE course_id = $1
ORDER BY matirial_id
LIMIT $2
OFFSET $3;

-- name: DeleteMatirials :exec
DELETE FROM matirials
WHERE matirial_id = $1 AND course_id = $2;