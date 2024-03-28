-- name: ListEnrolments :many
SELECT * FROM course_enrolments
WHERE student_id = $1 AND course_id = $2
ORDER BY enrolment_id
LIMIT $3
OFFSET $4;