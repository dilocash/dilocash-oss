-- Create "profiles" table
CREATE TABLE "public"."profiles" (
  "user_id" uuid NOT NULL,
  "display_name" text NULL,
  "email" text NOT NULL,
  "accepted_terms_version" text NULL,
  "accepted_terms_at" timestamptz NULL,
  "allow_data_analysis" boolean NULL DEFAULT false,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("user_id"),
  CONSTRAINT "profiles_email_key" UNIQUE ("email"),
  CONSTRAINT "profiles_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "auth"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "commands" table
CREATE TABLE "public"."commands" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "command_status" integer NOT NULL,
  "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted" boolean NULL DEFAULT false,
  "profile_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "commands_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "public"."profiles" ("user_id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_commands_profile_id" to table: "commands"
CREATE INDEX "idx_commands_profile_id" ON "public"."commands" ("profile_id");
-- Create index "idx_commands_updated_at" to table: "commands"
CREATE INDEX "idx_commands_updated_at" ON "public"."commands" ("updated_at");
-- Create "intents" table
CREATE TABLE "public"."intents" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "text_message" text NULL,
  "audio_message" text NULL,
  "image_message" text NULL,
  "intent_status" integer NOT NULL,
  "requires_review" boolean NULL DEFAULT false,
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


-- Create "handle_new_user" function
CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO public.profiles (user_id, display_name, email)
  VALUES (NEW.id, NEW.raw_user_meta_data ->> 'display_name', NEW.email);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
-- Create trigger "on_auth_user_created"
CREATE TRIGGER on_auth_user_created
  AFTER INSERT ON auth.users
  FOR EACH ROW
  EXECUTE PROCEDURE public.handle_new_user();