CREATE TYPE "user_role" AS ENUM (
  'admin',
  'student',
  'teacher'
);

CREATE TYPE "type_resource" AS ENUM (
  'pdf',
  'video',
  'image',
  'doc'
);

CREATE TABLE "users" (
  "user_name" varchar PRIMARY KEY,
  "role" user_role NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "is_email_verified" bool NOT NULL DEFAULT false,
  "password_changed_at" timestamp NOT NULL DEFAULT '0001-01-01',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "admins" (
  "admin_id" bigserial PRIMARY KEY,
  "user_name" varchar
);

CREATE TABLE "teachers" (
  "teacher_id" bigserial PRIMARY KEY,
  "admin_id" bigint,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "user_name" varchar,
  "hashed_password" varchar NOT NULL,
  "is_active" bool NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "students" (
  "student_id" bigserial PRIMARY KEY,
  "user_name" varchar
);

CREATE TABLE "courses" (
  "course_id" bigserial PRIMARY KEY,
  "teacher_id" bigint,
  "title" varchar NOT NULL,
  "type" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "course_progress" (
  "courseprogress_id" bigserial PRIMARY KEY,
  "enrolment_id" bigint,
  "progress" varchar NOT NULL
);

CREATE TABLE "course_enrolments" (
  "enrolment_id" bigserial PRIMARY KEY,
  "course_id" bigint,
  "request_id" bigint,
  "student_id" bigint
);

CREATE TABLE "assignments" (
  "assignment_id" bigserial PRIMARY KEY,
  "course_id" bigint,
  "type" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "submission_date" date NOT NULL
);

CREATE TABLE "submissions" (
  "submission_id" bigserial PRIMARY KEY,
  "assignment_id" bigint,
  "student_id" bigint
);

CREATE TABLE "requests" (
  "request_id" bigserial PRIMARY KEY,
  "student_id" bigint,
  "teacher_id" bigint,
  "course_id" bigint,
  "is_active" bool,
  "is_pending" bool,
  "is_accepted" bool,
  "is_declined" bool
);

CREATE TABLE "resources" (
  "resource_id" bigserial PRIMARY KEY,
  "course_id" bigint,
  "assignment_id" bigint,
  "title" varchar NOT NULL,
  "type" type_resource NOT NULL,
  "content_url" varchar NOT NULL
);

CREATE INDEX ON "admins" ("user_name");

CREATE UNIQUE INDEX ON "admins" ("user_name");

CREATE INDEX ON "teachers" ("admin_id");

CREATE UNIQUE INDEX ON "teachers" ("user_name");

CREATE UNIQUE INDEX ON "students" ("user_name");

CREATE INDEX ON "courses" ("teacher_id");

CREATE INDEX ON "course_progress" ("courseprogress_id");

CREATE INDEX ON "course_enrolments" ("course_id");

CREATE INDEX ON "course_enrolments" ("student_id");

CREATE INDEX ON "course_enrolments" ("course_id", "student_id");

CREATE INDEX ON "assignments" ("course_id");

CREATE INDEX ON "submissions" ("assignment_id");

CREATE INDEX ON "submissions" ("student_id");

CREATE INDEX ON "requests" ("student_id");

CREATE INDEX ON "requests" ("teacher_id");

CREATE INDEX ON "requests" ("course_id");

CREATE INDEX ON "requests" ("student_id", "teacher_id", "course_id");

CREATE INDEX ON "resources" ("course_id");

CREATE INDEX ON "resources" ("assignment_id");

CREATE INDEX ON "resources" ("course_id", "assignment_id");

ALTER TABLE "admins" ADD FOREIGN KEY ("user_name") REFERENCES "users" ("user_name");

ALTER TABLE "teachers" ADD FOREIGN KEY ("admin_id") REFERENCES "admins" ("admin_id");

ALTER TABLE "teachers" ADD FOREIGN KEY ("user_name") REFERENCES "users" ("user_name");

ALTER TABLE "students" ADD FOREIGN KEY ("user_name") REFERENCES "users" ("user_name");

ALTER TABLE "courses" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("teacher_id");

ALTER TABLE "course_progress" ADD FOREIGN KEY ("enrolment_id") REFERENCES "course_enrolments" ("enrolment_id");

ALTER TABLE "course_enrolments" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id");

ALTER TABLE "course_enrolments" ADD FOREIGN KEY ("request_id") REFERENCES "requests" ("request_id");

ALTER TABLE "course_enrolments" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("student_id");

ALTER TABLE "assignments" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id");

ALTER TABLE "submissions" ADD FOREIGN KEY ("assignment_id") REFERENCES "assignments" ("assignment_id");

ALTER TABLE "submissions" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("student_id");

ALTER TABLE "requests" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("student_id");

ALTER TABLE "requests" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("teacher_id");

ALTER TABLE "requests" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id");

ALTER TABLE "resources" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id");

ALTER TABLE "resources" ADD FOREIGN KEY ("assignment_id") REFERENCES "assignments" ("assignment_id");
