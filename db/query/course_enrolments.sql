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
WHERE course_id = $1 
ORDER BY enrolment_id
LIMIT $2
OFFSET $3;

-- name: ListEnrolmentsByUser :many
SELECT * FROM course_enrolments
WHERE user_id = $1 
ORDER BY enrolment_id
LIMIT $2
OFFSET $3;