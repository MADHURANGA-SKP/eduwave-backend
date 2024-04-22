-- name: CreateRequest :one
INSERT INTO requests (
    user_id,
    course_id,
    is_active,
    is_pending,
    is_accepted,
    is_declined
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetRequest :one
SELECT * FROM requests
WHERE user_id = $1;

-- name: UpdateRequests :one
UPDATE requests
SET is_active = $3, is_pending = $4, is_accepted = $5, is_declined = $6 
WHERE user_id = $1 AND course_id = $2
RETURNING *;

-- name: ListRequest :many
SELECT * FROM requests
ORDER BY request_id
LIMIT $1
OFFSET $2;

-- name: ListRequestByUser :many
SELECT * FROM requests
WHERE user_id = $1
ORDER BY request_id
LIMIT $2
OFFSET $3;

-- name: ListRequestByCourse :many
SELECT * FROM requests
WHERE course_id = $1
ORDER BY request_id
LIMIT $2
OFFSET $3;

-- name: DeleteRequest :exec
DELETE FROM requests
WHERE request_id = $1;