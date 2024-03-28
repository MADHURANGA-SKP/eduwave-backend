-- name: CreateAssignment :one
INSERT INTO assignments (
    type,
    title,
    description,
    submittion_date
) VALUES (
    $1, $2, $3, $4
)  RETURNING *;

-- name: GetAssignment :one
SELECT * FROM assignments
WHERE course_id = $1;

-- name: UpdateAssignment :one
UPDATE assignments
SET type = $2, title = $3, description = $4, submittion_date = $5
WHERE course_id = $1
RETURNING *;

-- name: DeleteAssignment :exec
DELETE FROM assignments
WHERE assignment_id = $1;