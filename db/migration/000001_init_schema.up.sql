CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "customer_id" bigint,
  "status" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "modified_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "license" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "modified_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "bookings" (
  "id" bigserial PRIMARY KEY,
  "customer_id" bigint,
  "flight_id" bigint,
  "booking_code" varchar NOT NULL,
  "status" varchar NOT NULL,
  "booked_date" timestamptz NOT NULL DEFAULT 'now()',
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "modified_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "flights" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "from" varchar NOT NULL,
  "to" varchar NOT NULL,
  "status" varchar NOT NULL,
  "available_slot" bigint NOT NULL,
  "departure_date" date NOT NULL,
  "arrival_date" date NOT NULL,
  "departure_time" varchar NOT NULL,
  "arrival_time" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "modified_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "users" ("customer_id");

CREATE INDEX ON "customers" ("id");

CREATE INDEX ON "bookings" ("customer_id");

CREATE INDEX ON "bookings" ("flight_id");

CREATE UNIQUE INDEX ON "bookings" ("flight_id", "customer_id");

CREATE INDEX ON "flights" ("id");

COMMENT ON COLUMN "flights"."available_slot" IS 'must be positive';

ALTER TABLE "users" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "bookings" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "bookings" ADD FOREIGN KEY ("flight_id") REFERENCES "flights" ("id");
