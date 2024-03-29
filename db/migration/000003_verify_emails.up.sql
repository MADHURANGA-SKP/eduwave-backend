CREATE TABLE "verify_emails" (
  "email_id" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("user_name") REFERENCES "users" ("user_name");