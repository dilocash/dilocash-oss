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
-- Create "commands" table
CREATE TABLE "public"."commands" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "status" text NOT NULL,
  "category" text NULL DEFAULT 'uncategorized',
  "description" text NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted" boolean NULL DEFAULT false,
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "commands_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_commands_last_updated_at" to table: "commands"
CREATE INDEX "idx_commands_last_updated_at" ON "public"."commands" ("updated_at");
-- Create index "idx_commands_user_id" to table: "commands"
CREATE INDEX "idx_commands_user_id" ON "public"."commands" ("user_id");
-- Create "intents" table
CREATE TABLE "public"."intents" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "text_message" text NULL,
  "audio_message" text NULL,
  "image_message" text NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted" boolean NULL DEFAULT false,
  "command_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "intents_command_id_fkey" FOREIGN KEY ("command_id") REFERENCES "public"."commands" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "transactions" table
CREATE TABLE "public"."transactions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "amount" numeric(19,4) NOT NULL,
  "currency" character(3) NOT NULL,
  "category" text NULL DEFAULT 'uncategorized',
  "description" text NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted" boolean NULL DEFAULT false,
  "command_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "transactions_command_id_fkey" FOREIGN KEY ("command_id") REFERENCES "public"."commands" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
