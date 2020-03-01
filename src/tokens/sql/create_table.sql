
-- Create
CREATE TABLE IF NOT EXISTS tokens (
    -- pk
    token CHAR(64) PRIMARY KEY,

    -- properties

    -- times
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);

-- Drop
DROP TABLE IF EXISTS accounts;

-- Trigger --
CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS 
$ $
BEGIN
    NEW.updated_at = NOW();

    RETURN NEW;
END;
$ $ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp 
    BEFORE UPDATE ON accounts 
    FOR EACH ROW 
    EXECUTE PROCEDURE trigger_set_timestamp();