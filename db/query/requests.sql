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
WHERE request_id = $1;

-- name: UpdateRequests :one
UPDATE requests
SET is_active = $2, is_pending = $3, is_accepted = $4, is_declined = $5 
WHERE user_id = $1
RETURNING *;

-- name: ListRequest :many
SELECT * FROM requests
ORDER BY request_id
LIMIT $1
OFFSET $2;

-- name: DeleteRequest :exec
DELETE FROM requests
WHERE request_id = $1;