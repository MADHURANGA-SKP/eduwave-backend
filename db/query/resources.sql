--name CreateResource :one
INSERT INTO students (
    user_name
) VALUES (
    $1
) RETURNING *;

--name: GetResource :one
SELECT * FROM students
WHERE student_id = $1;

--name: UpdateResource :one
UPDATE students
SET user_name = $2
WHERE student_id = $1
RETURNING *;

--name: ListResource :many
SELECT * FROM students
WHERE user_name = $1
ORDER BY student_id
LIMIT $2
OFFSET $3;

--name: DeleteResource :exec
DELETE FROM students
WHERE student_id = $1;