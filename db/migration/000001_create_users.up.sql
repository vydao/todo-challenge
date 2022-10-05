CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" char NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);