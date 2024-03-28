-- name: CreateStudent :one
INSERT INTO students (
    user_name
) VALUES (
    $1
) RETURNING *;

-- name: GetStudent :one
SELECT * FROM students
WHERE student_id = $1;

-- name: UpdateStudent :one
UPDATE students
SET user_name = $2
WHERE student_id = $1
RETURNING *;

-- name: ListStudents :many
SELECT * FROM students
WHERE user_name = $1
ORDER BY student_id
LIMIT $2
OFFSET $3;

-- name: DeleteStudent :exec
DELETE FROM students
WHERE student_id = $1;