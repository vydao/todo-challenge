CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "score" float8 NOT NULL DEFAULT 0.0,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);