-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "email" text NOT NULL,
  "accepted_terms_version" text NULL,
  "accepted_terms_at" timestamptz NULL,
  "allow_data_analysis" boolean NULL DEFAULT false,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "users_email_key" UNIQUE ("email")
);
-- Create "transactions" table
CREATE TABLE "public"."transactions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "amount" numeric(19,4) NOT NULL,
  "currency" character(3) NOT NULL,
  "category" text NULL DEFAULT 'uncategorized',
  "description" text NULL,
  "raw_input" text NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "transactions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_transactions_created_at" to table: "transactions"
CREATE INDEX "idx_transactions_created_at" ON "public"."transactions" ("created_at");
-- Create index "idx_transactions_user_id" to table: "transactions"
CREATE INDEX "idx_transactions_user_id" ON "public"."transactions" ("user_id");
