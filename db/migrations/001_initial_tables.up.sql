CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE IF NOT EXISTS journal (
    id UUID DEFAULT public.uuid_generate_v4() PRIMARY KEY,
    journal_name TEXT NOT NULL,
    UNIQUE(journal_name)
);

CREATE TABLE IF NOT EXISTS entry (
    id UUID DEFAULT public.uuid_generate_v4() PRIMARY KEY,
    journal_id UUID NOT NULL,
    content TEXT NOT NULL,
    CONSTRAINT fk_journal
        FOREIGN KEY(journal_id)
            REFERENCES journal(id)
);