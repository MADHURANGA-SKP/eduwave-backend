-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
    user_name,
    email,
    secret_code
) VALUES (
    $1, $2, $3
) RETURNING *;

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
