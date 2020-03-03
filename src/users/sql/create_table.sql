
-- Create
CREATE TABLE IF NOT EXISTS users (
    -- pk
    user_id CHAR(32) PRIMARY KEY,

    -- properties
    username VARCHAR(256) NOT NULL UNIQUE,
    password CHAR(64) NOT NULL,

    -- fk
    -- role

    -- times
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO users(
    user_id, username, password
) VALUES (
    'db780439d285e8aba7bf64daba277ec8',
    'kayac',
    '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92'
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