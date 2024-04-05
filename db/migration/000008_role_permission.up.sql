CREATE TABLE "role_permission" (
  "role_id" bigint,
  "permission_id" bigint,
  PRIMARY KEY ("role_id", "permission_id")
);

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("permission_id");

