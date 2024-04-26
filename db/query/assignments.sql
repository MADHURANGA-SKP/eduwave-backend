-- name: CreateAssignment :one
INSERT INTO assignments (
    resource_id,
    type,
    title,
    description,
    submission_date
) VALUES (
    $1, $2, $3, $4, $5
)  RETURNING *;

-- name: GetAssignment :one
SELECT * FROM assignments
WHERE assignment_id = $1;

-- name: GetAssignmentByResource :one
SELECT * FROM assignments
WHERE resource_id = $1;

-- name: UpdateAssignment :one
UPDATE assignments
SET type = $2, title = $3, description = $4, submission_date = $5
WHERE assignment_id = $1
RETURNING *;

-- name: DeleteAssignment :exec
DELETE FROM assignments
WHERE assignment_id = $1;