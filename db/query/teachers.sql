-- name: CreateTeacher :one
INSERT INTO teachers(
    user_id,
    admin_id,
    full_name,
    email,
    qualification,
    user_name,
    hashed_password,
    is_active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetTeacher :one
SELECT * FROM teachers
WHERE teacher_id = $1;

-- name: UpdateTeacher :one
UPDATE teachers
SET full_name = $2, email = $3, user_name = $4, hashed_password = $5, is_active = $6
WHERE teacher_id = $1
RETURNING *;

-- name: ListTeacher :many
SELECT * FROM teachers
WHERE teacher_id = $1
ORDER BY teacher_id
LIMIT $2
OFFSET $3;

-- name: DeleteTeacher :exec
DELETE FROM teachers
WHERE teacher_id = $1;

