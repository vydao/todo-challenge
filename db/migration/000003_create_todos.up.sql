CREATE TABLE "todos" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "challenge_id" bigint NOT NULL,
  "period" varchar NOT NULL,
  "point" float8 NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);