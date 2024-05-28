CREATE TABLE "user_space_activities" (
  "id" VARCHAR(26),
  "user_space_id" VARCHAR(26),
  "type" VARCHAR(100) NOT NULL,
  "data" JSONB NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);
