CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "urls" (
  "id" bigserial PRIMARY KEY,
  "user_id" integer NOT NULL,
  "org_url" varchar NOT NULL,
  "short_url" varchar UNIQUE NOT NULL,
  "clicks" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "urls" ("short_url");

CREATE INDEX ON "urls" ("user_id");

ALTER TABLE "urls" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
