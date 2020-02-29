
CREATE TABLE IF NOT EXISTS games (
    -- pk
    game_id CHAR(32) PRIMARY KEY,

    -- properties
    name VARCHAR(256) NOT NULL UNIQUE,
    href VARCHAR NOT NULL UNIQUE,
    category VARCHAR(32) NOT NULL,
    state CHAR(1) NOT NULL,

    -- times
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS games;

-- Trigger --
CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS 
$ $
BEGIN
    NEW.updated_at = NOW();

    RETURN NEW;
END;
$ $ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp 
    BEFORE UPDATE ON games 
    FOR EACH ROW 
    EXECUTE PROCEDURE trigger_set_timestamp();