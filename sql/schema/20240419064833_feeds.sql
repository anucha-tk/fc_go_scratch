-- +goose Up
CREATE TABLE "feeds" (
  "id" uuid PRIMARY KEY,
  "name" text NOT NULL,
  "url" text UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "user_id" uuid NOT NULL
);

ALTER TABLE "feeds" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

-- +goose Down
DROP TABLE feeds;
