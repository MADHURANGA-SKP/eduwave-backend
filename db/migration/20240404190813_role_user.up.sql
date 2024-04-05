CREATE TABLE "user_roles" (
  "user_id" bigserial,
  "role_id" bigint,
  PRIMARY KEY ("user_id", "role_id")
);

ALTER TABLE "user_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "user_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "teachers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
