--name: CreateAssignment :one
INSERT INTO assignments (
    type,
    title,
    description,
) VALUES (
    $1, $2, $3
)

--name: GetAssignment :one
SELECT * FROM assignments
WHERE course_id = $1;

--name: UpdateAssignment :one
UPDATE assignments
SET type = $2, title = $3, description = $4, 
WHERE course_id = $1
RETURNING *;

--name: DeleteAssignment :
DELETE FROM assignments
WHERE assignment_id = $1;