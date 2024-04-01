-- name: CreateRequest :one
INSERT INTO requests (
    is_active,
    is_pending,
    is_accepted,
    is_declined
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetRequest :one
SELECT * FROM requests
WHERE student_id = $1 AND request_id = $2;

-- name: UpdateRequests :one
UPDATE requests
SET is_active = $4, is_pending = $5, is_accepted = $6, is_declined = $7 
WHERE student_id = $1 AND teacher_id = $2 AND course_id =$3
RETURNING *;

-- name: ListRequest :many
SELECT * FROM requests
WHERE student_id = $1 AND teacher_id = $2 AND course_id =$3
ORDER BY request_id
LIMIT $4
OFFSET $5;

-- name: DeleteRequest :exec
DELETE FROM requests
WHERE student_id = $1 AND request_id = $2;