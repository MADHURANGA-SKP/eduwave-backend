

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("permission_id");

CREATE TABLE "teachers_users" (
  "teachers_teacher_id" bigserial,
  "users_user_id" bigint,
  PRIMARY KEY ("teachers_teacher_id", "users_user_id")
);

ALTER TABLE "teachers_users" ADD FOREIGN KEY ("teachers_teacher_id") REFERENCES "teachers" ("teacher_id");

ALTER TABLE "teachers_users" ADD FOREIGN KEY ("users_user_id") REFERENCES "users" ("user_id");

-- ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");
