CREATE TABLE "users" (
  "id" VARCHAR(26) PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL UNIQUE,
  "password_hash" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);
