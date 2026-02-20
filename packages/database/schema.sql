-- Copyright (c) 2026 dilocash
-- Use of this source code is governed by an MIT-style
-- license that can be found in the LICENSE file.

-- Users Table: Handles the "Consent Gate" and Training Opt-in
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    accepted_terms_version TEXT,
    accepted_terms_at TIMESTAMP WITH TIME ZONE,
    allow_data_analysis BOOLEAN DEFAULT FALSE, -- ADR-026: Opt-in only
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Commands Table: The actions requested by the user
CREATE TABLE IF NOT EXISTS commands (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    command_status INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN DEFAULT FALSE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
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
CREATE INDEX IF NOT EXISTS idx_commands_user_id ON commands(user_id);
CREATE INDEX IF NOT EXISTS idx_commands_last_updated_at ON commands(updated_at);