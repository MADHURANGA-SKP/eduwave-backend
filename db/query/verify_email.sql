-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
    email,
    secret_code
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetVerifyEmail :one
SELECT * FROM verify_emails
WHERE secret_code = $1 LIMIT 1;

-- name: UpdateVerifyEmail :one
UPDATE verify_emails
SET
    is_used = TRUE
WHERE
    email_id = @email_id
    AND secret_code = @secret_code
    AND is_used = FALSE
    AND expired_at > now()
RETURNING *;
