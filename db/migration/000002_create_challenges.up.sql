CREATE TABLE "challenges" (
  "id" BIGSERIAL PRIMARY KEY,
  "start_date" timestamp NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);