-- name: CreateEnrolments :one
INSERT INTO course_enrolments (
    course_id,
    request_id,
    user_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ListEnrolments :many
SELECT * FROM course_enrolments
ORDER BY enrolment_id
LIMIT $1
OFFSET $2;