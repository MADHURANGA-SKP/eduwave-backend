-- name: CreateCourses :one
INSERT INTO courses (
    user_id,
    title,
    type,
    description,
    image
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetCourses :one
SELECT * FROM courses
WHERE course_id = $1;

-- name: GetCourseByUserCourse :one
SELECT * FROM courses
WHERE user_id = $1 AND course_id = $2
LIMIT 1;

-- name: UpdateCourses :one
UPDATE courses
SET title = $2, type = $3, description = $4, image = $5
WHERE course_id = $1
RETURNING *;

-- name: ListCourses :many
SELECT * FROM courses
ORDER BY course_id
LIMIT $1
OFFSET $2;

-- name: ListCoursesByUser :many
SELECT * FROM courses
WHERE user_id = $1
ORDER BY course_id
LIMIT $2
OFFSET $3;

-- name: DeleteCourses :exec
DELETE FROM courses
WHERE course_id = $1;