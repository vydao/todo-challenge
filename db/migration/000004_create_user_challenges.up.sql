CREATE TABLE "competitions" (
  "id" BIGSERIAL PRIMARY KEY,
  "challenger_id" bigint NOT NULL,
  "rival_id" bigint NOT NULL,
  "challenge_id" bigint NOT NULL,
  "status" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);