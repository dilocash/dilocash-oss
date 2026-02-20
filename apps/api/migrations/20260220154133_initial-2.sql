-- Modify "commands" table
ALTER TABLE "public"."commands" DROP COLUMN "status", DROP COLUMN "category", DROP COLUMN "description", ADD COLUMN "command_status" integer NOT NULL;
-- Modify "intents" table
ALTER TABLE "public"."intents" ADD COLUMN "intent_status" integer NOT NULL, ADD COLUMN "requires_review" boolean NULL DEFAULT false;
