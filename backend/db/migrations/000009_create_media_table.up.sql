CREATE TABLE "media" (
  "id" VARCHAR(26) PRIMARY KEY,
  "user_id" VARCHAR(26),
  "user_space_id" VARCHAR(26),
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);
