-- +goose Up
CREATE TABLE "feed_follows" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "user_id" uuid NOT NULL,
  "feed_id" uuid NOT NULL
);

CREATE UNIQUE INDEX ON "feed_follows" ("user_id", "feed_id");

ALTER TABLE "feed_follows" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "feed_follows" ADD FOREIGN KEY ("feed_id") REFERENCES "feeds" ("id") ON DELETE CASCADE;

-- +goose Down
DROP TABLE feed_follows;
