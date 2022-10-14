CREATE TABLE "competition_todos" (
  "id" BIGSERIAL PRIMARY KEY,
  "competition_id" bigint NOT NULL,
  "todo_id" bigint NOT NULL,
  "is_completed" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);