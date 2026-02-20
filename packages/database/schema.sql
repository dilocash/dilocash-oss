-- Copyright (c) 2026 dilocash
-- Use of this source code is governed by an MIT-style
-- license that can be found in the LICENSE file.
CREATE SCHEMA IF NOT EXISTS auth;
CREATE TABLE IF NOT EXISTS auth.users (id UUID PRIMARY KEY, email TEXT, raw_user_meta_data JSONB);

-- Users Table: Handles the "Consent Gate" and Training Opt-in
CREATE TABLE IF NOT EXISTS public.profiles (
    id UUID REFERENCES auth.users NOT NULL PRIMARY KEY,
    display_name TEXT NULL,
    email TEXT NULL,
    accepted_terms_version TEXT,
    accepted_terms_at TIMESTAMP WITH TIME ZONE,
    allow_data_analysis BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE public.profiles ENABLE ROW LEVEL SECURITY;

-- function to create a user when a new user is created via supabase signup
CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO public.profiles (id, display_name, email)
  VALUES (NEW.id, NEW.raw_user_meta_data ->> 'display_name', NEW.email);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

-- Trigger to create a profile when a new user is created via supabase signup
CREATE TRIGGER on_auth_user_created
  AFTER INSERT ON auth.users
  FOR EACH ROW
  EXECUTE PROCEDURE public.handle_new_user();

-- Commands Table: The actions requested by the user
CREATE TABLE IF NOT EXISTS commands (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    command_status INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN DEFAULT FALSE,
    profile_id UUID NOT NULL REFERENCES public.profiles(id) ON DELETE CASCADE
);

-- Intents Table: The actions taken by the user
CREATE TABLE IF NOT EXISTS intents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    text_message TEXT,
    audio_message TEXT,
    image_message TEXT,
    intent_status INTEGER NOT NULL,
    requires_review BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN DEFAULT FALSE,
    command_id UUID NOT NULL REFERENCES commands(id) ON DELETE CASCADE
);

-- Transactions Table: The actions taken by the user
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- Using NUMERIC for precision: 19 total digits, 4 after decimal
    amount NUMERIC(19, 4) NOT NULL,
    currency CHAR(3) NOT NULL, -- ISO 4217 (USD, EUR, etc.)
    category TEXT DEFAULT 'uncategorized',
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN DEFAULT FALSE,
    command_id UUID NOT NULL REFERENCES commands(id) ON DELETE CASCADE
);

-- Indexing for performance
CREATE INDEX IF NOT EXISTS idx_commands_profile_id ON commands(profile_id);
CREATE INDEX IF NOT EXISTS idx_commands_updated_at ON commands(updated_at);
