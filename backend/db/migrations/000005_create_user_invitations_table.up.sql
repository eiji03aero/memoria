CREATE TABLE "user_invitations" (
  "id" VARCHAR(26) PRIMARY KEY,
  "user_id" VARCHAR(26) NOT NULL,
  "type" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL,

  CONSTRAINT fk_user_invitations_users
    FOREIGN KEY(user_id)
    REFERENCES users(id)
);
