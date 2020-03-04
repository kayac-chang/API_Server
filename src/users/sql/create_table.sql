
-- Create
CREATE TABLE IF NOT EXISTS users (
    -- pk
    user_id CHAR(32) PRIMARY KEY,

    -- properties
    username VARCHAR(256) NOT NULL UNIQUE,

    -- fk
    -- role

    -- times
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Test
INSERT INTO users(
    user_id, username
) VALUES (
    'db780439d285e8aba7bf64daba277ec8',
    'kayac'
);

-- Drop
DROP TABLE IF EXISTS users;

-- Trigger --
CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS 
$$
BEGIN
    NEW.updated_at = NOW();

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp 
    BEFORE UPDATE ON users 
    FOR EACH ROW 
    EXECUTE PROCEDURE trigger_set_timestamp();