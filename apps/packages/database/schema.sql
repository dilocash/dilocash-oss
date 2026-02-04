-- TODO add license header

-- Users Table: Handles the "Consent Gate" and Training Opt-in
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    accepted_terms_version TEXT,
    accepted_terms_at TIMESTAMP WITH TIME ZONE,
    allow_data_analysis BOOLEAN DEFAULT FALSE, -- ADR-026: Opt-in only
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Transactions Table: The core ledger
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- Using NUMERIC for precision: 19 total digits, 4 after decimal
    amount NUMERIC(19, 4) NOT NULL,
    currency CHAR(3) NOT NULL, -- ISO 4217 (USD, EUR, etc.)
    
    category TEXT DEFAULT 'uncategorized',
    description TEXT,
    raw_input TEXT,            -- Stores the original transcript for auditing
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexing for performance
CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_transactions_created_at ON transactions(created_at);