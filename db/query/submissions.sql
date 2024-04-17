-- name: CreateSubmission :one
INSERT INTO submissions (
    assignment_id,
    user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetsubmissionsByAssignment :one
SELECT * FROM submissions
WHERE assignment_id = $1 
LIMIT 1;

-- name: GetsubmissionsByUser :one
SELECT * FROM submissions
WHERE user_id = $1
LIMIT 1;

-- name: Listsubmissions :many
SELECT * FROM submissions
ORDER BY submission_id
LIMIT $1
OFFSET $2;