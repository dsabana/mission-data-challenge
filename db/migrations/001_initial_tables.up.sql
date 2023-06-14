CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE IF NOT EXISTS journal (
    id UUID DEFAULT public.uuid_generate_v4() PRIMARY KEY,
    journal_name TEXT NOT NULL,
    UNIQUE(journal_name)
);