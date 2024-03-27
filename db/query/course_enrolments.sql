--name: ListEnrolments :many
SELECT * FROM course_enrolments
WHERE student_id = $1 AND request_id = $2 AND course_id =$3
ORDER BY enrolment_id
LIMIT $4
OFFSET $5;