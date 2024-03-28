-- name: GetCourseProgress :one
SELECT * FROM course_progress
WHERE courseprogress_id = $1 AND enrolment_id = $2 
LIMIT 1;

-- name: ListCourseProgress :many
SELECT * FROM course_progress
WHERE enrolment_id = $1
ORDER BY courseprogress_id
LIMIT $2
OFFSET $3;