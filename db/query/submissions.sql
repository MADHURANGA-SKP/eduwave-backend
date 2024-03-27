--name Getsubmissions :one
SELECT * FROM submissions
WHERE assignment_id = $1 AND student_id = $2 
LIMIT 1;

--name: Listsubmissions :many
SELECT * FROM submissions
WHERE assignment_id = $1 AND student_id = $2
ORDER BY submission_id
LIMIT $2
OFFSET $3;